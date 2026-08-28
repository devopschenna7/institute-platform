// const API_URL = "http://localhost:8080/api/v1";

const API_URL = "/api/v1";

export async function register(
  studentCode,
  email,
  password
) {
  const response = await fetch(
    `${API_URL}/auth/register`,
    {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        student_code: studentCode,
        email: email,
        password: password,
      }),
    }
  );

  const data = await response.json().catch(() => null);

  if (!response.ok) {
    throw new Error(
      typeof data === "string"
        ? data
        : data?.message || "Registration failed"
    );
  }

  //sandeep

  if (!response.ok) {
  const errorMessage =
    typeof data === "string"
      ? data
      : data?.message || "Registration failed";

  if (
    errorMessage.includes("students_student_code_key") ||
    errorMessage.includes("duplicate key") ||
    errorMessage.includes("already exists")
  ) {
    throw new Error(
      "Student already exists. Please use a different student code."
    );
  }

  throw new Error(errorMessage);
}

  return data;
}

// export async function createRegistration(studentId, courseId) {
//   const response = await fetch(`${API_URL}/registrations`, {
//     method: "POST",
//     credentials: "include",
//     headers: {
//       "Content-Type": "application/json",
//     },
//     body: JSON.stringify({
//       student_id: studentId,
//       course_id: courseId,
//     }),
//   });

//   const data = await response.json().catch(() => null);

//   if (!response.ok) {
//     throw new Error(
//       typeof data === "string"
//         ? data
//         : data?.message || "Registration failed"
//     );
//   }

//   return data;
// }

export async function createRegistration(studentId, courseId) {
  const response = await fetch(`${API_URL}/registrations`, {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      student_id: studentId,
      course_id: courseId,
    }),
  });

  const text = await response.text();

  if (!response.ok) {
    throw new Error(
      text || "Registration failed"
    );
  }

  try {
    return JSON.parse(text);
  } catch {
    return text;
  }
}

export async function getMyRegistrations(studentId) {
  const response = await fetch(
    `${API_URL}/students/${studentId}/registrations`,
    {
      credentials: "include",
    }
  );

  const data = await response.json().catch(() => null);

  if (!response.ok) {
    throw new Error(
      typeof data === "string"
        ? data
        : data?.message || "Failed to fetch registrations"
    );
  }

  return data;
}

export async function getCourses() {
  const response = await fetch(`${API_URL}/courses`);

  if (!response.ok) {
    throw new Error("Failed to fetch courses");
  }

  return response.json();
}

export async function getCourseBySlug(slug) {
  const response = await fetch(`${API_URL}/courses/${slug}`);

  if (!response.ok) {
    throw new Error("Course not found");
  }

  return response.json();
}
//sandeep
export async function getDashboard() {
  const response = await fetch(
    `${API_URL}/students/me/dashboard`,
    {
      credentials: "include",
    }
  );

  const data = await response.json().catch(() => null);

  if (!response.ok) {
    throw new Error(
      typeof data === "string"
        ? data
        : data?.message || "Failed to get dashboard"
    );
  }

  return data;
}

export async function getMe() {
  const response = await fetch(`${API_URL}/auth/me`, {
    credentials: "include",
  });

  if (!response.ok) {
    throw new Error("Not authenticated");
  }

  return response.json();
}

export async function logout() {
  const response = await fetch(`${API_URL}/auth/logout`, {
    method: "POST",
    credentials: "include",
  });

  if (!response.ok) {
    throw new Error("Logout failed");
  }

  return response.json();
}

export async function login(email, password) {
  const response = await fetch(`${API_URL}/auth/login`, {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email,
      password,
    }),
  });

  const data = await response.json().catch(() => null);

  if (!response.ok) {
    throw new Error(
      typeof data === "string"
        ? data
        : data?.message || "Login failed"
    );
  }

  return data;
}