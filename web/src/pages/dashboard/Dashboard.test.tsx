import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { vi, describe, it, expect, beforeEach } from 'vitest'
import { BrowserRouter } from 'react-router-dom'
import Dashboard from './Dashboard'

// Mock the auth context
const mockAuthContext = {
  user: {
    id: '1',
    name: 'John Doe',
    email: 'john@example.com',
    role: 'employee',
  },
  isAuthenticated: true,
  logout: vi.fn(),
  loading: false,
}

vi.mock('@/contexts/AuthContext', () => ({
  useAuth: () => mockAuthContext,
}))

// Mock the course service
const mockGetCourses = vi.fn()
const mockGetUserProgress = vi.fn()

vi.mock('@/services/courseService', () => ({
  getCourses: mockGetCourses,
  getUserProgress: mockGetUserProgress,
}))

// Mock the analytics service
const mockGetAnalytics = vi.fn()

vi.mock('@/services/analyticsService', () => ({
  getAnalytics: mockGetAnalytics,
}))

const renderDashboard = () => {
  return render(
    <BrowserRouter>
      <Dashboard />
    </BrowserRouter>
  )
}

describe('Dashboard Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    
    // Default mock responses
    mockGetCourses.mockResolvedValue([
      {
        id: '1',
        title: 'Introduction to Go',
        description: 'Learn the basics of Go programming',
        category: 'Programming',
        progress: 75,
      },
      {
        id: '2',
        title: 'Advanced React',
        description: 'Advanced React concepts and patterns',
        category: 'Frontend',
        progress: 30,
      },
    ])
    
    mockGetUserProgress.mockResolvedValue({
      totalCourses: 2,
      completedCourses: 1,
      inProgressCourses: 1,
      averageProgress: 52.5,
    })
    
    mockGetAnalytics.mockResolvedValue({
      weeklyProgress: [20, 35, 45, 60, 75, 80, 85],
      topCategories: ['Programming', 'Frontend', 'Backend'],
      timeSpent: 120,
    })
  })

  it('renders dashboard with user information', () => {
    renderDashboard()
    
    expect(screen.getByText(/welcome back, john doe/i)).toBeInTheDocument()
    expect(screen.getByText(/john@example.com/i)).toBeInTheDocument()
    expect(screen.getByText(/employee/i)).toBeInTheDocument()
  })

  it('displays course progress cards', async () => {
    renderDashboard()
    
    await waitFor(() => {
      expect(screen.getByText(/introduction to go/i)).toBeInTheDocument()
      expect(screen.getByText(/advanced react/i)).toBeInTheDocument()
      expect(screen.getByText(/75%/i)).toBeInTheDocument()
      expect(screen.getByText(/30%/i)).toBeInTheDocument()
    })
  })

  it('shows progress statistics', async () => {
    renderDashboard()
    
    await waitFor(() => {
      expect(screen.getByText(/total courses/i)).toBeInTheDocument()
      expect(screen.getByText(/2/i)).toBeInTheDocument()
      expect(screen.getByText(/completed/i)).toBeInTheDocument()
      expect(screen.getByText(/1/i)).toBeInTheDocument()
      expect(screen.getByText(/in progress/i)).toBeInTheDocument()
      expect(screen.getByText(/52.5%/i)).toBeInTheDocument()
    })
  })

  it('displays analytics data', async () => {
    renderDashboard()
    
    await waitFor(() => {
      expect(screen.getByText(/weekly progress/i)).toBeInTheDocument()
      expect(screen.getByText(/top categories/i)).toBeInTheDocument()
      expect(screen.getByText(/time spent/i)).toBeInTheDocument()
      expect(screen.getByText(/2h/i)).toBeInTheDocument()
    })
  })

  it('handles loading state', () => {
    mockGetCourses.mockImplementation(() => new Promise(resolve => setTimeout(resolve, 100)))
    mockGetUserProgress.mockImplementation(() => new Promise(resolve => setTimeout(resolve, 100)))
    mockGetAnalytics.mockImplementation(() => new Promise(resolve => setTimeout(resolve, 100)))
    
    renderDashboard()
    
    expect(screen.getByText(/loading/i)).toBeInTheDocument()
  })

  it('handles error state for courses', async () => {
    mockGetCourses.mockRejectedValueOnce(new Error('Failed to fetch courses'))
    
    renderDashboard()
    
    await waitFor(() => {
      expect(screen.getByText(/failed to load courses/i)).toBeInTheDocument()
    })
  })

  it('handles error state for progress', async () => {
    mockGetUserProgress.mockRejectedValueOnce(new Error('Failed to fetch progress'))
    
    renderDashboard()
    
    await waitFor(() => {
      expect(screen.getByText(/failed to load progress/i)).toBeInTheDocument()
    })
  })

  it('handles error state for analytics', async () => {
    mockGetAnalytics.mockRejectedValueOnce(new Error('Failed to fetch analytics'))
    
    renderDashboard()
    
    await waitFor(() => {
      expect(screen.getByText(/failed to load analytics/i)).toBeInTheDocument()
    })
  })

  it('navigates to course detail when course card is clicked', async () => {
    const mockNavigate = vi.fn()
    vi.mocked(require('react-router-dom')).useNavigate = () => mockNavigate
    
    renderDashboard()
    
    await waitFor(() => {
      const courseCard = screen.getByText(/introduction to go/i).closest('div')
      if (courseCard) {
        fireEvent.click(courseCard)
        expect(mockNavigate).toHaveBeenCalledWith('/courses/1')
      }
    })
  })

  it('shows empty state when no courses are available', async () => {
    mockGetCourses.mockResolvedValueOnce([])
    
    renderDashboard()
    
    await waitFor(() => {
      expect(screen.getByText(/no courses available/i)).toBeInTheDocument()
      expect(screen.getByText(/start exploring our course catalog/i)).toBeInTheDocument()
    })
  })

  it('displays course categories correctly', async () => {
    renderDashboard()
    
    await waitFor(() => {
      expect(screen.getByText(/programming/i)).toBeInTheDocument()
      expect(screen.getByText(/frontend/i)).toBeInTheDocument()
    })
  })

  it('shows progress bar for each course', async () => {
    renderDashboard()
    
    await waitFor(() => {
      const progressBars = screen.getAllByRole('progressbar')
      expect(progressBars).toHaveLength(2)
    })
  })

  it('handles logout when logout button is clicked', () => {
    renderDashboard()
    
    const logoutButton = screen.getByRole('button', { name: /logout/i })
    fireEvent.click(logoutButton)
    
    expect(mockAuthContext.logout).toHaveBeenCalled()
  })

  it('shows user avatar or initials', () => {
    renderDashboard()
    
    // Check for user initials (JD for John Doe)
    expect(screen.getByText(/jd/i)).toBeInTheDocument()
  })

  it('displays recent activity section', async () => {
    renderDashboard()
    
    await waitFor(() => {
      expect(screen.getByText(/recent activity/i)).toBeInTheDocument()
    })
  })

  it('shows quick actions', () => {
    renderDashboard()
    
    expect(screen.getByText(/browse courses/i)).toBeInTheDocument()
    expect(screen.getByText(/view progress/i)).toBeInTheDocument()
    expect(screen.getByText(/take quiz/i)).toBeInTheDocument()
  })

  it('handles responsive design for mobile', () => {
    // Mock window.innerWidth for mobile
    Object.defineProperty(window, 'innerWidth', {
      writable: true,
      configurable: true,
      value: 375,
    })
    
    renderDashboard()
    
    // Should still render all content
    expect(screen.getByText(/welcome back, john doe/i)).toBeInTheDocument()
  })

  it('updates progress in real-time when data changes', async () => {
    const { rerender } = renderDashboard()
    
    await waitFor(() => {
      expect(screen.getByText(/75%/i)).toBeInTheDocument()
    })
    
    // Update mock data
    mockGetCourses.mockResolvedValueOnce([
      {
        id: '1',
        title: 'Introduction to Go',
        description: 'Learn the basics of Go programming',
        category: 'Programming',
        progress: 90,
      },
    ])
    
    // Re-render with new data
    rerender(
      <BrowserRouter>
        <Dashboard />
      </BrowserRouter>
    )
    
    await waitFor(() => {
      expect(screen.getByText(/90%/i)).toBeInTheDocument()
    })
  })

  it('handles network timeout gracefully', async () => {
    mockGetCourses.mockImplementation(() => 
      new Promise((_, reject) => setTimeout(() => reject(new Error('Timeout')), 100))
    )
    
    renderDashboard()
    
    await waitFor(() => {
      expect(screen.getByText(/timeout/i)).toBeInTheDocument()
    })
  })

  it('shows appropriate message for different user roles', () => {
    // Test with admin role
    mockAuthContext.user.role = 'admin'
    
    renderDashboard()
    
    expect(screen.getByText(/admin dashboard/i)).toBeInTheDocument()
  })

  it('handles empty analytics data gracefully', async () => {
    mockGetAnalytics.mockResolvedValueOnce({
      weeklyProgress: [],
      topCategories: [],
      timeSpent: 0,
    })
    
    renderDashboard()
    
    await waitFor(() => {
      expect(screen.getByText(/no analytics data available/i)).toBeInTheDocument()
    })
  })
})
