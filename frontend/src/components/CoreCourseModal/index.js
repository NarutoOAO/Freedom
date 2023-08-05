import React, { useState } from 'react'
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import './style.css'
export default function CoreCourseModel() {
  // State to manage the visibility of the modal
  const [show, setShow] = useState(false);
  // Get the token from sessionStorage
  const token = sessionStorage.getItem('token');
  // Use to get core course inforamtion
  const [corecourses, setCoreCourses] = useState(null);
  // Function to handle showing the modal and fetch core course infroamtion from server
  const handleShow = async () => {
    setShow(true);
    const response = await fetch('http://127.0.0.1:5005/api/v1/user_mandatory_course', {
      method: 'POST',
      headers: {
        'Authorization': token,
      },
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      setCoreCourses(data.data);
    }
  }

  // Function to handle closing the modal
  const handleClose = () => {
    setShow(false);
  }

 
  return (
    <>
      <button onClick={handleShow}>Look</button>
      <Modal show={show} onHide={handleClose} size='lg'>
        <Modal.Header closeButton>
          <Modal.Title>Handbook(For core course)</Modal.Title>
        </Modal.Header>
        <Modal.Body className='modalBody' style={{ maxHeight: '500px', overflowY: 'auto' }}>
          <table className="table table-hover">
            <thead>
              <tr>
                <th scope="col">#</th>
                <th scope="col">Course Number</th>
                <th scope="col">Course Name</th>
                <th scope="col">Term</th>
                <th scope="col">Max People</th>
                <th scope="col">Credit</th>
              </tr>
            </thead>
            <tbody>
              {/* Map through core course information and display them in a table */}
              {corecourses !== null && corecourses !== '' && corecourses !== undefined &&
                corecourses.map((courseItem,index) => (
                  <tr>
                    <th scope="row">{index+1}</th>
                    <td>{courseItem.CourseNumber}</td>
                    <td>{courseItem.CourseName}</td>
                    <td>{courseItem.Term}</td>
                    <td>{courseItem.MaxPeople}</td>
                    <td>6 uoc</td>
                  </tr>
                ))}
            </tbody>
          </table>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="primary" onClick={handleClose}>close</Button>
        </Modal.Footer>
      </Modal>
    </>
  )
}

