import React, { useState } from 'react'
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import './style.css'
// define the modal to select the course
export default function SelectCourseModal(props) {
  const [show, setShow] = useState(false);
  const [selectedCourse, setSelectedCourse] = useState('');
  const token = sessionStorage.getItem('token');
  const [courses, setCourses] = useState(null);
  // Open the modal to select the course
  const handleShow = async () => {
    setShow(true);
    // get the courses
    const response = await fetch('http://127.0.0.1:5005/api/v1/student-select-course', {
      method: 'GET',
      headers: {
        'Authorization': token,
      },
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      console.log(data.data);
      setCourses(data.data);
    }
  }
// close the modal
  const handleClose = () => {
    setCourses(null);
    setSelectedCourse('');
    setShow(false);
  }
// define the function to select a course
  const handleEnrollCourse = async () => {
    const apiUrl = 'http://127.0.0.1:5005/api/v1/student-course';
    try {
      const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
          'Content-type': 'application/json',
          'Authorization': token,
        },
        body: JSON.stringify({ course_number: parseInt(selectedCourse) }),
      });

      const data = await response.json();
      if (data.status===200) {
        alert('Course enrollment successful!');

      } else {
        alert(data.msg);
      }
    } catch (error) {
      alert('Failed to enroll in the course. Please try again.');
      console.error(error);
    }
    setSelectedCourse('');
    handleShow();
    props.getCoursesFn();
  };
// define the function to select a course
  const handleSlectCourse =(course_number)=>{
    setSelectedCourse(course_number);
    //console.log(course_number);
  }
  return (
    <>
      <button onClick={handleShow}>Select</button>
      <Modal show={show} onHide={handleClose} size='lg'>
        <Modal.Header closeButton>
          <Modal.Title>Course List</Modal.Title>
        </Modal.Header>
        <Modal.Body className='modalBody' style={{ maxHeight: '500px', overflowY: 'auto' }}>
          <table className="table table-hover">
          {/* <colgroup>
              <col style={{ width: '1%' }} />
              <col style={{ width: '10%' }} />
              <col style={{ width: '15%' }} />
              <col style={{ width: '70%' }} />
              <col style={{ width: '10%' }} />
              <col style={{ width: '15%' }} />
              <col style={{ width: '70%' }} />
            </colgroup> */}
            <thead>
              <tr>
                <th scope="col">#</th>
                <th scope="col">Course Number</th>
                <th scope="col">Course Name</th>
                <th scope="col">Teacher Name</th>
                <th scope="col">CLass Time</th>
                <th scope="col">Current People</th>
                <th scope="col">Max People</th>
                <th scope="col">Type</th>
              </tr>
            </thead>
            <tbody>
              {courses !== null && courses !== '' && courses !== undefined &&
                courses.map((courseItem,index) => (
                  <tr key={index} onClick={() => handleSlectCourse(courseItem.CourseNumber)}
                  className={selectedCourse === courseItem.CourseNumber ? 'selectedCourseEnroll' : ''}>
                    <th scope="row">{index+1}</th>
                    <td>{courseItem.CourseNumber}</td>
                    <td>{courseItem.CourseName}</td>
                    <td>{courseItem.TeacherName}</td>
                    <td>{courseItem.ClassTime}</td>
                    <td>{courseItem.CurrentPeople}</td>
                    <td>{courseItem.MaxPeople}</td>
                    <td>{courseItem.Classification}</td>
                  </tr>
                ))}
            </tbody>
          </table>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="primary" onClick={handleEnrollCourse}>Enroll</Button>
        </Modal.Footer>
      </Modal>
    </>
  )
}

