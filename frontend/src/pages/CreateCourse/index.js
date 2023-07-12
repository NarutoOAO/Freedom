import React, { useState } from 'react';
import './style.css'
function CreateCourse() {
  const [courseNumber, setCourseNumber] = useState('');
  const [courseName, setCourseName] = useState('');
  const token = localStorage.getItem('token');
  const handleCreateCourse = async () => {
    const response = await fetch('http://127.0.0.1:5005/api/v1/teacher-course', {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
        'Authorization': token,
      },
      body: JSON.stringify({
        course_number: parseInt(courseNumber),
        course_name: courseName
      })
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      // console.log(data);
      alert("Succeed!");
    }
    setCourseName('');
    setCourseNumber('');
  };

  return (
    <div className='CreateCourse'>
      <h2>Create course</h2>
      <form>
        <label>
          Course Number:
          <input
            type="text"
            value={courseNumber}
            onChange={(e) => setCourseNumber(e.target.value)}
          />
        </label>
        <br />
        <label>
        Course Name:
          <input
            type="text"
            value={courseName}
            onChange={(e) => setCourseName(e.target.value)}
          />
        </label>
        <br />
        <button type="button" onClick={handleCreateCourse}>
          Create
        </button>
      </form>
    </div>
  );
}

export default CreateCourse;
