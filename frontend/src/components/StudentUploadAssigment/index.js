import Modal from 'react-bootstrap/Modal';
import React, { useState } from 'react';
import { Button } from "antd";

function StudentUploadAssigment(props) {
    
    const [show, setShow] = useState(false);
    // Open the modal
    const handleAsigmentUploadShow = () => setShow(true);
    // Close the modal and reset input values
    const handleAsigmentUploaClose = () => {
        setShow(false);
    };
    function handleStudentSolutionUpload(props) {
        const fileInput = document.getElementById('fileInput');
        const file = fileInput.files[0];
        const studentSumbitSolution = new FormData();
        const token = sessionStorage.getItem('token');
        const currentDateTime = new Date();
        if (new Date(props.end_time) < currentDateTime) {
          alert('Missed submission deadline.');
          return;
        }
        
        studentSumbitSolution.append('assignment_id', parseInt(props.assignment_id));
        studentSumbitSolution.append('file', file);
        // Validate file
        if (!file ) {
          alert('Please submit a file');
          return;
        }
    
        // Validate file type
        if (file.type !== 'application/pdf') {
          alert('Only PDF files are allowed.');
          return;
        }
       
        // Send the assignment solution to the server
        fetch('http://127.0.0.1:5005/api/v1/assignment_solution', {
          method: 'POST',
          headers: {
            'Authorization': token,
          },
          body:studentSumbitSolution,
        })
          .then(response => response.json())
          .then((data) => {
            if (data.status !== 200) {
              alert(data.msg);
            } else {
              alert('Succeed!');
              handleAsigmentUploaClose();
            }
          }) 
          .catch((error) => {
            alert("Failed!")
          });
          
      }

    return (
            <div className='createAssigment'>
              <button
                variant="primary"
                onClick={handleAsigmentUploadShow}
                className="assigment-function-button" 
                style={{ backgroundColor: '#cfd8dc' }}
              >
                Upload Solution
              </button>
              
              <Modal show={show} onHide={handleAsigmentUploaClose}>
                <Modal.Header closeButton>
                  <Modal.Title>Upload Solution</Modal.Title>
                </Modal.Header>
                <Modal.Body className="modalBody">
                  <form encType="multipart/form-data">
                    <div>
                      <label htmlFor="fileInput">File:</label><br />
                      <input type="file" id="fileInput" />
                      <div style={{ color: 'gray', fontStyle: 'italic' ,marginTop:'10px'}}>Only PDF files are allowed</div>
                    </div>
                    
                  </form>
                </Modal.Body>
                <Modal.Footer>
                  <Button variant="secondary" onClick={handleAsigmentUploaClose}>
                      Close
                  </Button>
                  <Button variant="primary" onClick={() => handleStudentSolutionUpload(props)}>
                        Upload
                  </Button>
                </Modal.Footer>
              </Modal>
            </div>
       
    )

}
export default StudentUploadAssigment;