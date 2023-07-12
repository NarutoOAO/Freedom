import React, { useState } from 'react';
import './style.css'
function EnrollCourse() {
  const [courseNumber, setCourseNumber] = useState('');
  const token = localStorage.getItem('token');
  const handleEnrollCourse = async () => {
    const apiUrl = 'http://127.0.0.1:5005/api/v1/student-course';

    try {
      const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
          'Content-type': 'application/json',
          'Authorization': token,
        },
        body: JSON.stringify({ course_number: parseInt(courseNumber) }),
      });

      const data = await response.json();
      console.log(data);
      if (data.status===200) {
        // Enrollment successful
        // console.log(data);
        alert('Course enrollment successful!');
      } else {
        // Enrollment failed
        alert(data.msg);
      }
    } catch (error) {
      // Error occurred during enrollment
      alert('Failed to enroll in the course. Please try again.');
      console.error(error);
    }

    // Reset the course number input
    setCourseNumber('');
  };

  return (
    <div className="enroll-course-container">
      <h2>Enroll in a Course</h2>
      <input
        type="text"
        placeholder="Course Number"
        value={courseNumber}
        onChange={(e) => setCourseNumber(e.target.value)}
      />
      <button onClick={handleEnrollCourse}>Enroll</button>
    </div>
  );
}

export default EnrollCourse;

