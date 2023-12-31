import React, { useEffect, useState } from 'react';
import Form from 'react-bootstrap/Form';
import SelectCourseModal from '../../components/SelectCourseModal';
import CheckStudentStatus from '../../components/CheckStudentStatus';
import CoreCourseModal from '../../components/CoreCourseModal';
import './style.css'
// define the page for student to enroll course
function EnrollCourse() {
  // Get the token from session storage 
  const token = sessionStorage.getItem('token');
  // State to store the selected course for dropping
  const [selectedCourse, setSelectedCourse] = useState('');
   // State to store the list of selected courses
  const [selectedCourses, setSelectedCourses] = useState('');

  useEffect (()=>{
    getCourses();
    // eslint-disable-next-line
  }, []);

  // Fetch the list of enrolled courses
  const getCourses = async () => {
    const res = await fetch('http://127.0.0.1:5005/api/v1/student-course',{
      method: 'GET',
      headers: {
        'Authorization': token,
      },
    });
    const data=await res.json();
    if(data.status===200){
      setSelectedCourses(data.data);
    }else{
      alert(data.msg);
    }
  }
  
  // Handle the course drop action
  const DropCourse = async () =>{
    const res = await fetch('http://127.0.0.1:5005/api/v1/student_course/'+selectedCourse,{
      method: "DELETE",
      headers:{
        "Authorization": token,
      }
    })
    const data = await res.json();
    if(data.status===200){
      alert("Drop succeed!");
      getCourses();
    }else{
      alert("Drop failed!");
    }
  }

  useEffect(() => {
     // Reset the selected course when the list of enrolled courses changes
    if (Array.isArray(selectedCourses) && selectedCourses.length > 0) {
      setSelectedCourse(selectedCourses[0].CourseNumber);
    }
  }, [selectedCourses]);
  
  return (
    <>
    {/* Check My Status Sectionm, jump to components/CheckStudentStatus*/}
    <div className="enroll-course-container">
      <h2>Check My Status</h2>
      <CheckStudentStatus/>
    </div>
     {/*Enroll course Sectionm, jump to components/SelectCourseModal*/}
    <div className="enroll-course-container">
      <h2>Enroll Course</h2>
      <SelectCourseModal getCoursesFn={getCourses}/>
    </div>
     {/*Handbook Sectionm, jump to components/CoreCourseModal*/}
    <div className="enroll-course-container">
      <h2>Handbook(For core course)</h2>
      <CoreCourseModal getCoursesFn={getCourses}/>
     {/* Drop Course Section */}
    </div>
        <div className="enroll-course-container">
        <h2>Drop Course</h2>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Label>Your Courses</Form.Label>
              <Form.Select
                value={selectedCourse}
                onChange={(event) => setSelectedCourse(event.target.value)}
              >
                {selectedCourses !== null && selectedCourses !== '' ? (
                  selectedCourses.map((course, index) => (
                    <option key={index} value={course.CourseNumber}>
                      {course.CourseNumber} {course.CourseName}
                    </option>
                  ))
                ) : (
                  <option disabled>You haven't enrolled any course!</option>
                )}
              </Form.Select>
            </Form.Group>
        <button onClick={DropCourse}>Drop</button>
      </div>
      </>
  );
}

export default EnrollCourse;

