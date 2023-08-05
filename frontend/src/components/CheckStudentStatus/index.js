import React, { useState ,useEffect} from 'react'
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import './style.css'
export default function CheckStudentStatus(props) {
  // State to manage the visibility of the modal
  const [show, setShow] = useState(false);
  // Get the token from sessionStorage
  const token = sessionStorage.getItem('token');
  // Get the studyOption from sessionStorage
  const studyoption = sessionStorage.getItem('studyOption');
  // Use to store adk score
  const [adkCourseCredit,setAdkCourseCredit]=useState('0');
  // Use to store de score
  const [deCourseCredit,deAdkCourseCredit]=useState('0');
  // Use to store core course score
  const [coreCourseCredit,coreAdkCourseCredit]=useState('0');
  // Function to fetch credit information from the server
  const getCredit = async () => {
    const dataInformation= {
    };
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/student-course-statistics' , {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
        body: JSON.stringify(dataInformation),
      });

      if (response.status === 200) {
        const responseData = await response.json();
        setAdkCourseCredit(responseData.data[0].Credit);
        deAdkCourseCredit(responseData.data[1].Credit);
        coreAdkCourseCredit(responseData.data[2].Credit);
      } else {
        throw new Error('Failed to Publish');
      }
    } catch (error) {
      console.error(error);
    }
  };
  
  useEffect(() => {
    getCredit();
    // eslint-disable-next-line
  },[]);

  // Function to handle showing the modal
  const handleShow = async () => {
    getCredit();
    setShow(true);
  }

  // Function to handle closing the modal
  const handleClose = () => {
    setShow(false);
  }

  return (
    <>
      <button onClick={handleShow}>Check</button>
      <Modal show={show} onHide={handleClose} size='lg'>
        <Modal.Header closeButton>
          <Modal.Title>You Status</Modal.Title>
        </Modal.Header>
        <Modal.Body className='modalBody' style={{  fontFamily: 'Arial, sans-serif', fontSize: '18px', color: '#333', backgroundColor: '#f9f9f9', padding: '10px', border: '1px solid #ccc', borderRadius: '5px', boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)',maxHeight: '500px',overflowY: 'auto', }}>
            {/* Display credit */}
            <p>You have completed: </p>
            <p>Advanced Disciplinary Knowledge(ADK): {adkCourseCredit} uoc</p>
            <p>Disciplinary Electives(DE): {deCourseCredit} uoc</p> 
            <p style={{marginBottom:'40px'}}>Core Courses: {coreCourseCredit} uoc</p>
            {/* Display additional information based on the study option */}
            {studyoption === 'AI' ? (
            <div>
                <p style={{marginBottom:'40px'}}>As a {studyoption} student, you need to finish 24 uoc for adk, 30 uoc for de, and 42 uoc for core to graduate.</p>
                <p>Course list that you must finish before graduating:</p>
                <p>Comp 9021, Comp 9024, Comp9032, Comp9311, Comp9417, Comp9900, comp9414</p>
            </div>
            ) : (
            <div>
                <p style={{marginBottom:'40px'}}>As an {studyoption} student, you need to finish 18 uoc for adk, 42 uoc for de, and 36 uoc for core to graduate.</p>
                <p>Course list that you must finish before graduating:</p>
                <p>Comp 9021, Comp 9024, Comp9032, Comp9311, Comp9900, comp9020</p>
            </div>
            )}
        </Modal.Body>
        <Modal.Footer>
          <Button variant="primary" onClick={handleClose}>Close</Button>
        </Modal.Footer>
      </Modal>
    </>
  )
}

