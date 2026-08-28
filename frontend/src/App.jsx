import { useEffect, useState } from "react";

import {
  BrowserRouter,
  Routes,
  Route,
  Link,
  useNavigate,
  useParams,
  useLocation,
} from "react-router-dom";

import instructorPhoto from "./assets/instructor.png";

import {
  register,
  getCourses,
  getCourseBySlug,
  login,
  getMe,
  getDashboard,
  logout,
  createRegistration,
  getMyRegistrations,
} from "./api/api";

/* =========================================================
   Protected Route
========================================================= */

function ProtectedRoute({ children }) {
  const [loading, setLoading] = useState(true);
  const [authenticated, setAuthenticated] = useState(false);

  useEffect(() => {
    getMe()
      .then(() => {
        setAuthenticated(true);
      })
      .catch(() => {
        window.location.href = "/login";
      })
      .finally(() => {
        setLoading(false);
      });
  }, []);

  if (loading) {
    return (
      <div className="page-loader">
        <div className="loader"></div>
        <p>Checking your account...</p>
      </div>
    );
  }

  if (!authenticated) {
    return null;
  }

  return children;
}

/* =========================================================
   Logged-In Header
========================================================= */

function Header() {
  const navigate = useNavigate();

  async function handleLogout() {
    try {
      await logout();

      navigate("/", {
        replace: true,
      });
    } catch (err) {
      console.error("Logout failed:", err);
    }
  }

  return (
    <header className="header">
      <div className="header-inner">

        <Link
          to="/courses"
          className="logo"
        >
          CODE<span>KERDOS</span>
        </Link>

        <nav className="nav">

          <Link to="/courses">
            Courses
          </Link>

          

          <Link to="/my-registrations">
            My Registrations
          </Link>

          <button
            className="nav-logout"
            onClick={handleLogout}
          >
            Logout
          </button>

        </nav>

      </div>
    </header>
  );
}

/* =========================================================
   Public Header
========================================================= */

function PublicHeader() {
  return (
    <header className="landing-header">

      <Link
        to="/"
        className="logo"
      >
        CODE<span>KERDOS</span>
      </Link>

      <nav className="landing-nav">

        <a href="/#courses">
          Courses
        </a>

        <a href="/#instructor">
          Instructor
        </a>

      </nav>

      <div className="landing-actions">

        <Link
          to="/login"
          className="btn btn-outline"
        >
          Login
        </Link>

        <Link
          to="/register"
          className="btn btn-primary"
        >
          Register
        </Link>

      </div>

    </header>
  );
}

/* =========================================================
   Landing Page
========================================================= */

function Home() {
  const [courses, setCourses] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    async function loadCourses() {
      try {
        const data = await getCourses();

        setCourses(data);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    }

    loadCourses();
  }, []);

  return (
    <div className="landing-page">

      <header className="landing-header">

        <Link
          to="/"
          className="logo"
        >
          CODE<span>KERDOS</span>
        </Link>

        <nav className="landing-nav">

          <a href="#courses">
            Courses
          </a>

          <a href="#instructor">
            Instructor
          </a>

        </nav>

        <div className="landing-actions">

          <Link
            to="/login"
            className="btn btn-outline"
          >
            Login
          </Link>

          <Link
            to="/register"
            className="btn btn-primary"
          >
            Register
          </Link>

        </div>

      </header>

      <main>

        {/* =================================================
            Hero
        ================================================= */}

        <section className="hero-section">

          <div className="hero-content">

            <div className="hero-badge">
              Learn. Build. Grow.
            </div>

            <h1>
              Build your skills.
              <br />
              Shape your future.
            </h1>

            <p>
              Learn practical technology skills through
              industry-focused courses designed to help
              you build, grow and succeed.
            </p>

            <div className="hero-buttons">

              <a
                href="#courses"
                className="btn btn-primary btn-large"
              >
                Explore Courses
              </a>

              <Link
                to="/register"
                className="btn btn-outline btn-large"
              >
                Create Student Account
              </Link>

            </div>

          </div>

          <div className="hero-card">

            <div className="hero-card-icon">
              {"</>"}
            </div>

            <h3>
              Practical Learning
            </h3>

            <p>
              Learn technologies through focused courses
              and real-world skills.
            </p>

          </div>

        </section>

        {/* =================================================
            Courses
        ================================================= */}

        <section
          id="courses"
          className="landing-section courses-section"
        >

          <div className="landing-section-heading">

            <div>

              <p className="eyebrow">
                AVAILABLE COURSES
              </p>

              <h2>
                Explore Our Courses
              </h2>

              <p>
                Browse available courses without creating
                an account. Login is required only when
                you want to register.
              </p>

            </div>

            <span className="course-count">

              {loading
                ? "Loading..."
                : `${courses.length} Courses`}

            </span>

          </div>

          {error && (
            <div className="error-message">
              {error}
            </div>
          )}

          {!loading &&
            !error &&
            courses.length === 0 && (
              <div className="empty-state">

                <h3>
                  No courses available
                </h3>

                <p>
                  There are currently no courses
                  available.
                </p>

              </div>
            )}

          <div className="course-grid landing-course-grid">

            {courses.map((course) => (

              <div
                className="course-card"
                key={course.id}
              >

                <div className="course-card-top">

                  <span className="course-level">
                    {course.level}
                  </span>

                  <span className="course-duration">
                    {course.duration}
                  </span>

                </div>

                <h3>
                  {course.name}
                </h3>

                <p>
                  {course.description}
                </p>

                <Link
                  to={`/courses/${course.slug}`}
                  className="course-link"
                >
                  View Course
                  <span>→</span>
                </Link>

              </div>

            ))}

          </div>

        </section>

        {/* =================================================
            Instructor
        ================================================= */}

        <section
          id="instructor"
          className="instructor-section"
        >

          <div className="instructor-photo-wrap">

            <img
              src={instructorPhoto}
              alt="Instructor"
              className="instructor-photo"
            />

          </div>

          <div className="instructor-content">

            <p className="eyebrow">
              MEET YOUR INSTRUCTOR
            </p>

            <h2>
              Chenna Ageswar Rao
            </h2>

            <h3>
              DevOps & Cloud Engineering Professional
            </h3>

            <p>
              Learn practical DevOps and cloud technologies
              from an industry professional with hands-on
              experience in modern infrastructure and
              cloud-native technologies.
            </p>

            <div className="instructor-skills">

              <span>Kubernetes</span>
              <span>Docker</span>
              <span>Jenkins</span>
              <span>AWS</span>

            </div>

            <a
              href="https://www.linkedin.com/in/chenna-ageswar-rao-45956025a/"
              target="_blank"
              rel="noreferrer"
              className="btn btn-primary"
            >
              View LinkedIn Profile
            </a>

          </div>

        </section>

        {/* =================================================
            Features
        ================================================= */}

        <section className="features landing-features">

          <div className="feature">

            <div className="feature-icon">
              01
            </div>

            <h3>
              Learn
            </h3>

            <p>
              Build a strong foundation with practical,
              focused learning.
            </p>

          </div>

          <div className="feature">

            <div className="feature-icon">
              02
            </div>

            <h3>
              Build
            </h3>

            <p>
              Turn your knowledge into real-world skills
              and projects.
            </p>

          </div>

          <div className="feature">

            <div className="feature-icon">
              03
            </div>

            <h3>
              Grow
            </h3>

            <p>
              Keep improving and move confidently toward
              your career goals.
            </p>

          </div>

        </section>

      </main>

      {/* =================================================
          Footer
      ================================================= */}

      <footer className="landing-footer">

        <div>

          <Link
            to="/"
            className="logo"
          >
            CODE<span>KERDOS</span>
          </Link>

          <p>
            Learn. Build. Grow.
          </p>

        </div>

        <div className="footer-links">

          <a href="#courses">
            Courses
          </a>

          <a href="#instructor">
            Instructor
          </a>

          <Link to="/login">
            Login
          </Link>

          <Link to="/register">
            Register
          </Link>

        </div>

      </footer>

    </div>
  );
}

/* =========================================================
   Root Route
   Redirect logged-in students to /courses
========================================================= */

function PublicHomeRoute() {
  const navigate = useNavigate();

  const [checkingAuth, setCheckingAuth] =
    useState(true);

  useEffect(() => {

    async function checkAuthentication() {

      try {

        await getMe();

        /*
         * Already logged in.
         * Never show the public landing page.
         */
        navigate("/courses", {
          replace: true,
        });

      } catch {

        /*
         * Not logged in.
         * Show public landing page.
         */

      } finally {

        setCheckingAuth(false);

      }

    }

    checkAuthentication();

  }, [navigate]);

  if (checkingAuth) {

    return (
      <div className="page-loader">

        <div className="loader"></div>

        <p>
          Loading...
        </p>

      </div>
    );

  }

  return <Home />;
}

/* =========================================================
   Register
========================================================= */

function Register() {
  const navigate = useNavigate();

  const [studentCode, setStudentCode] =
    useState("");

  const [email, setEmail] =
    useState("");

  const [password, setPassword] =
    useState("");

  const [error, setError] =
    useState("");

  const [loading, setLoading] =
    useState(false);

  const [success, setSuccess] =
    useState(false);

  async function handleRegister(event) {

    event.preventDefault();

    setError("");
    setLoading(true);

    try {

      await register(
        studentCode,
        email,
        password
      );

      setSuccess(true);

    } catch (err) {

      setError(err.message);

    } finally {

      setLoading(false);

    }

  }

  if (success) {

    return (
      <div className="auth-page">

        <div className="success-card">

          <div className="success-icon">
            ✓
          </div>

          <h1>
            Account Created!
          </h1>

          <p>
            Your student account has been created
            successfully.
          </p>

          <p>
            You can now login using your email and
            password.
          </p>

          <button
            className="btn btn-primary btn-full"
            onClick={() => navigate("/login")}
          >
            Go to Login
          </button>

        </div>

      </div>
    );
  }

  return (
    <div className="auth-page">

      <div className="auth-card">

        <Link
          to="/"
          className="auth-logo"
        >
          CODE<span>KERDOS</span>
        </Link>

        <div className="auth-heading">

          <h1>
            Create your account
          </h1>

          <p>
            Create your student account to access
            our courses.
          </p>

        </div>

        <form onSubmit={handleRegister}>

          <div className="form-group">

            <label>
              Student Name
            </label>

            <input
              id="studentCode"
              type="text"
              value={studentCode}
              onChange={(event) =>
                setStudentCode(
                  event.target.value
                )
              }
              placeholder="e.g. Chenna Rao"
              required
            />

          </div>

          <div className="form-group">

            <label>
              Email
            </label>

            <input
              type="email"
              value={email}
              onChange={(event) =>
                setEmail(
                  event.target.value
                )
              }
              placeholder="you@example.com"
              required
            />

          </div>

          <div className="form-group">

            <label>
              Password
            </label>

            <input
              type="password"
              value={password}
              onChange={(event) =>
                setPassword(
                  event.target.value
                )
              }
              placeholder="Create a password"
              required
            />

          </div>

          {error && (
            <div className="error-message">
              {error}
            </div>
          )}

          <button
            type="submit"
            className="btn btn-primary btn-full"
            disabled={loading}
          >
            {loading
              ? "Creating Account..."
              : "Create Account"}
          </button>

        </form>

        <p className="auth-footer">

          Already have an account?{" "}

          <Link to="/login">
            Login
          </Link>

        </p>

      </div>

    </div>
  );
}

/* =========================================================
   Login
========================================================= */

function Login() {
  const navigate = useNavigate();
  const location = useLocation();

  const [email, setEmail] =
    useState("");

  const [password, setPassword] =
    useState("");

  const [error, setError] =
    useState("");

  const [loading, setLoading] =
    useState(false);

  async function handleLogin(event) {

    event.preventDefault();

    setError("");
    setLoading(true);

    try {

      await login(
        email,
        password
      );

      navigate(
        location.state?.from ||
          "/courses",
        {
          replace: true,
        }
      );

    } catch (err) {

      setError(err.message);

    } finally {

      setLoading(false);

    }

  }

  return (
    <div className="auth-page">

      <div className="auth-card">

        <Link
          to="/"
          className="auth-logo"
        >
          CODE<span>KERDOS</span>
        </Link>

        <div className="auth-heading">

          <h1>
            Welcome back
          </h1>

          <p>
            Login to continue your learning journey.
          </p>

        </div>

        <form onSubmit={handleLogin}>

          <div className="form-group">

            <label>
              Email
            </label>

            <input
              type="email"
              value={email}
              onChange={(event) =>
                setEmail(
                  event.target.value
                )
              }
              placeholder="you@example.com"
              required
            />

          </div>

          <div className="form-group">

            <label>
              Password
            </label>

            <input
              type="password"
              value={password}
              onChange={(event) =>
                setPassword(
                  event.target.value
                )
              }
              placeholder="Enter your password"
              required
            />

          </div>

          {error && (
            <div className="error-message">
              {error}
            </div>
          )}

          <button
            type="submit"
            className="btn btn-primary btn-full"
            disabled={loading}
          >
            {loading
              ? "Logging in..."
              : "Login"}
          </button>

        </form>

        <p className="auth-footer">

          Don't have an account?{" "}

          <Link to="/register">
            Create one
          </Link>

        </p>

      </div>

    </div>
  );
}

/* =========================================================
   Courses / Student Home
========================================================= */

function Courses() {
  const [courses, setCourses] =
    useState([]);

  const [dashboard, setDashboard] =
    useState(null);

  const [error, setError] =
    useState("");

  const [loading, setLoading] =
    useState(true);

  useEffect(() => {

    async function loadData() {

      try {

        const [
          courseData,
          dashboardData,
        ] = await Promise.all([
          getCourses(),
          getDashboard(),
        ]);

        setCourses(courseData);
        setDashboard(dashboardData);

      } catch (err) {

        setError(err.message);

      } finally {

        setLoading(false);

      }

    }

    loadData();

  }, []);

  if (loading) {

    return (
      <div className="page-loader">

        <div className="loader"></div>

        <p>
          Loading courses...
        </p>

      </div>
    );

  }

  return (
    <div className="dashboard-page">

      <Header />

      <main className="dashboard-container">

        <section className="welcome-section">

          <div>

            <p className="eyebrow">
              STUDENT DASHBOARD
            </p>

            <h1>
              Welcome,{" "}
              {dashboard?.Student?.Name ||
                "Student"}{" "}
              👋
            </h1>

            <p>
              Explore our courses and find the right
              one for your learning journey.
            </p>

          </div>

        </section>

        <section className="courses-section">

          <div className="section-heading">

            <div>

              <p className="eyebrow">
                EXPLORE
              </p>

              <h2>
                Our Courses
              </h2>

            </div>

            <span className="course-count">
              {courses.length} Courses
            </span>

          </div>

          {error && (
            <div className="error-message">
              {error}
            </div>
          )}

          {!error &&
            courses.length === 0 && (
              <div className="empty-state">

                <h3>
                  No courses available
                </h3>

                <p>
                  There are currently no courses
                  available.
                </p>

              </div>
            )}

          <div className="course-grid">

            {courses.map((course) => (

              <div
                className="course-card"
                key={course.id}
              >

                <div className="course-card-top">

                  <span className="course-level">
                    {course.level}
                  </span>

                  <span className="course-duration">
                    {course.duration}
                  </span>

                </div>

                <h3>
                  {course.name}
                </h3>

                <p>
                  {course.description}
                </p>

                <Link
                  to={`/courses/${course.slug}`}
                  className="course-link"
                >
                  View Details
                  <span>→</span>
                </Link>

              </div>

            ))}

          </div>

        </section>

      </main>

    </div>
  );
}

/* =========================================================
   Course Details
========================================================= */

function CourseDetails() {
  const { slug } = useParams();
  const navigate = useNavigate();

  const [course, setCourse] =
    useState(null);

  const [studentId, setStudentId] =
    useState(null);

  const [error, setError] =
    useState("");

  const [registering, setRegistering] =
    useState(false);

  const [
    registrationSuccess,
    setRegistrationSuccess,
  ] = useState(false);

  const [
    registrationError,
    setRegistrationError,
  ] = useState("");

  /* =======================================================
     Load Course
  ======================================================= */

  useEffect(() => {

    async function loadCourse() {

      try {

        const data =
          await getCourseBySlug(slug);

        setCourse(data);

      } catch (err) {

        setError(err.message);

      }

    }

    loadCourse();

  }, [slug]);

  /* =======================================================
     Check Authentication
  ======================================================= */

  useEffect(() => {

    async function checkLogin() {

      try {

        const data =
          await getMe();

        setStudentId(
          data.student_id
        );

      } catch {

        setStudentId(null);

      }

    }

    checkLogin();

  }, []);

  /* =======================================================
     Register
  ======================================================= */

  async function handleRegister() {

    if (!course) {
      return;
    }

    /*
     * Not logged in
     */
    if (!studentId) {

      navigate("/login", {
        state: {
          from:
            `/courses/${course.slug}`,
        },
      });

      return;
    }

    setRegistering(true);
    setRegistrationError("");

    try {

      await createRegistration(
        studentId,
        course.id
      );

      setRegistrationSuccess(true);

    } catch (err) {

      setRegistrationError(
        err.message
      );

    } finally {

      setRegistering(false);

    }

  }

  /* =======================================================
     Error
  ======================================================= */

  if (error) {

    return (
      <div className="dashboard-page">

        {studentId ? (
          <Header />
        ) : (
          <PublicHeader />
        )}

        <main className="dashboard-container">

          <div className="error-page">

            <h1>
              Course not found
            </h1>

            <p>
              {error}
            </p>

            <button
              type="button"
              className="btn btn-primary"
              onClick={() => {

                if (studentId) {
                  navigate("/courses");
                } else {
                  navigate("/#courses");
                }

              }}
            >
              Back to Courses
            </button>

          </div>

        </main>

      </div>
    );
  }

  /* =======================================================
     Loading
  ======================================================= */

  if (!course) {

    return (
      <div className="page-loader">

        <div className="loader"></div>

        <p>
          Loading course...
        </p>

      </div>
    );

  }

  /* =======================================================
     Curriculum
  ======================================================= */

  const curriculum = [
    {
      number: "01",
      title: "Introduction & Fundamentals",
      description:
        "Understand the core concepts, terminology and fundamentals required to get started with this technology.",
    },
    {
      number: "02",
      title: "Core Concepts",
      description:
        "Learn the important concepts and components used in real-world projects.",
    },
    {
      number: "03",
      title: "Hands-on Implementation",
      description:
        "Apply your knowledge through practical exercises and real-world examples.",
    },
    {
      number: "04",
      title: "Advanced Concepts",
      description:
        "Explore advanced features, best practices and commonly used production patterns.",
    },
    {
      number: "05",
      title: "Real-World Project",
      description:
        "Build a practical project that brings together the concepts covered throughout the course.",
    },
  ];

  /* =======================================================
     Learning Points
  ======================================================= */

  const learningPoints = [
    "Understand the fundamentals and core concepts",
    "Work with the technology using practical examples",
    "Build and manage real-world applications",
    "Follow industry best practices",
    "Troubleshoot common problems",
    "Apply your knowledge through hands-on projects",
  ];

  return (
    <div className="dashboard-page">

      {/* =================================================
          Dynamic Header
      ================================================= */}

      {studentId ? (
        <Header />
      ) : (
        <PublicHeader />
      )}

      <main className="dashboard-container">

        {/* =================================================
            Back to Courses
        ================================================= */}

        <button
          type="button"
          className="back-link"
          onClick={() => {

            if (studentId) {
              navigate("/courses");
            } else {
              navigate("/#courses");
            }

          }}
        >
          ← Back to Courses
        </button>

        {/* =================================================
            Course Hero
        ================================================= */}

        <section className="course-detail-hero">

          <div className="course-detail-hero-content">

            <div className="course-detail-badges">

              <span className="course-level">
                {course.level}
              </span>

              <span className="course-duration">
                {course.duration}
              </span>

            </div>

            <p className="eyebrow">
              COURSE DETAILS
            </p>

            <h1>
              {course.name}
            </h1>

            <p className="course-detail-description">
              {course.description}
            </p>

            <div className="course-detail-meta">

              <div className="course-meta-item">

                <span className="meta-icon">
                  ◷
                </span>

                <div>

                  <small>
                    Duration
                  </small>

                  <strong>
                    {course.duration}
                  </strong>

                </div>

              </div>

              <div className="course-meta-item">

                <span className="meta-icon">
                  ◆
                </span>

                <div>

                  <small>
                    Level
                  </small>

                  <strong>
                    {course.level}
                  </strong>

                </div>

              </div>

              <div className="course-meta-item">

                <span className="meta-icon">
                  ✓
                </span>

                <div>

                  <small>
                    Learning
                  </small>

                  <strong>
                    Practical
                  </strong>

                </div>

              </div>

            </div>

          </div>

          {/* =================================================
              Registration Card
          ================================================= */}

          <div className="course-register-card">

            <div className="register-card-icon">
              {"</>"}
            </div>

            <h3>
              Ready to start learning?
            </h3>

            <p>
              Join this course and start building
              practical skills.
            </p>

            {!registrationSuccess ? (

              <>

                <button
                  className="btn btn-primary btn-large btn-full"
                  onClick={handleRegister}
                  disabled={registering}
                >
                  {registering
                    ? "Registering..."
                    : "Register for this Course"}
                </button>

                {!studentId && (

                  <p className="register-note">
                    Login is required to register
                    for this course.
                  </p>

                )}

                {registrationError && (

                  <div className="error-message">
                    {registrationError}
                  </div>

                )}

              </>

            ) : (

              <div className="registration-success">

                <div className="success-icon">
                  ✓
                </div>

                <div>

                  <h3>
                    Thanks for registering!
                  </h3>

                  <p>
                    Our team will connect with you
                    soon.
                  </p>

                </div>

              </div>

            )}

          </div>

        </section>

        {/* =================================================
            Curriculum
        ================================================= */}

        <section className="course-curriculum">

          <div className="course-section-heading">

            <p className="eyebrow">
              WHAT YOU'LL LEARN
            </p>

            <h2>
              The Curriculum
            </h2>

            <p>
              A practical learning path designed to
              take you from fundamentals to real-world
              implementation.
            </p>

          </div>

          <div className="curriculum-list">

            {curriculum.map((item) => (

              <div
                className="curriculum-item"
                key={item.number}
              >

                <div className="curriculum-number">
                  {item.number}
                </div>

                <div className="curriculum-content">

                  <h3>
                    {item.title}
                  </h3>

                  <p>
                    {item.description}
                  </p>

                </div>

                <div className="curriculum-arrow">
                  →
                </div>

              </div>

            ))}

          </div>

        </section>

        {/* =================================================
            What You'll Learn
        ================================================= */}

        <section className="what-you-learn">

          <div className="learn-content">

            <p className="eyebrow">
              COURSE OUTCOMES
            </p>

            <h2>
              What You'll Learn
            </h2>

            <p>
              By the end of this course, you will have
              practical knowledge that you can apply to
              real-world projects.
            </p>

          </div>

          <div className="learning-points">

            {learningPoints.map(
              (point, index) => (

                <div
                  className="learning-point"
                  key={index}
                >

                  <span className="learning-check">
                    ✓
                  </span>

                  <span>
                    {point}
                  </span>

                </div>

              )
            )}

          </div>

        </section>

        {/* =================================================
            Bottom CTA
        ================================================= */}

        <section className="course-bottom-cta">

          <div>

            <p className="eyebrow">
              START YOUR JOURNEY
            </p>

            <h2>
              Ready to build your skills?
            </h2>

            <p>
              Take the next step and start learning
              through practical, industry-focused
              training.
            </p>

          </div>

        </section>

      </main>

    </div>
  );
}

/* =========================================================
   My Registrations
========================================================= */

function MyRegistrations() {
  const [registrations, setRegistrations] =
    useState([]);

  const [courses, setCourses] =
    useState([]);

  const [error, setError] =
    useState("");

  const [loading, setLoading] =
    useState(true);

  useEffect(() => {

    async function loadRegistrations() {

      try {

        const me = await getMe();

        const [
          registrationData,
          courseData,
        ] = await Promise.all([
          getMyRegistrations(
            me.student_id
          ),
          getCourses(),
        ]);

        setRegistrations(
          registrationData
        );

        setCourses(courseData);

      } catch (err) {

        setError(err.message);

      } finally {

        setLoading(false);

      }

    }

    loadRegistrations();

  }, []);

  if (loading) {

    return (
      <div className="page-loader">

        <div className="loader"></div>

        <p>
          Loading your registrations...
        </p>

      </div>
    );

  }

  return (
    <div className="dashboard-page">

      <Header />

      <main className="dashboard-container">

        <section className="page-heading">

          <p className="eyebrow">
            STUDENT DASHBOARD
          </p>

          <h1>
            My Registrations
          </h1>

          <p>
            Courses you have registered for.
          </p>

        </section>

        {error && (
          <div className="error-message">
            {error}
          </div>
        )}

        {!error &&
          registrations.length === 0 && (

            <div className="empty-state">

              <div className="empty-icon">
                📚
              </div>

              <h3>
                No registrations yet
              </h3>

              <p>
                You haven't registered for any
                courses yet.
              </p>

              <Link
                to="/courses"
                className="btn btn-primary"
              >
                Explore Courses
              </Link>

            </div>

          )}

        <div className="registration-list">

          {registrations.map(
            (registration) => {

              const course =
                courses.find(
                  (course) =>
                    course.id ===
                    registration.course_id
                );

              return (

                <div
                  className="registration-card"
                  key={registration.id}
                >

                  <div>

                    <h2>
                      {course
                        ? course.name
                        : `Course ${registration.course_id}`}
                    </h2>

                    {course && (
                      <p>
                        {course.description}
                      </p>
                    )}

                    {course && (

                      <div className="registration-meta">

                        <span>
                          Duration:{" "}
                          {course.duration}
                        </span>

                        <span>
                          Level:{" "}
                          {course.level}
                        </span>

                      </div>

                    )}

                  </div>

                  <span
                    className={`status status-${String(
                      registration.status
                    ).toLowerCase()}`}
                  >
                    {registration.status}
                  </span>

                </div>

              );

            }
          )}

        </div>

      </main>

    </div>
  );
}

/* =========================================================
   App
========================================================= */

function App() {
  return (
    <BrowserRouter>

      <Routes>

        {/* Public / Root */}

        <Route
          path="/"
          element={
            <PublicHomeRoute />
          }
        />

        {/* Authentication */}

        <Route
          path="/login"
          element={<Login />}
        />

        <Route
          path="/register"
          element={<Register />}
        />

        {/* Student Home */}

        <Route
          path="/courses"
          element={
            <ProtectedRoute>
              <Courses />
            </ProtectedRoute>
          }
        />

        {/* Course Details */}

        <Route
          path="/courses/:slug"
          element={<CourseDetails />}
        />

        {/* My Registrations */}

        <Route
          path="/my-registrations"
          element={
            <ProtectedRoute>
              <MyRegistrations />
            </ProtectedRoute>
          }
        />

      </Routes>

    </BrowserRouter>
  );
}

export default App;