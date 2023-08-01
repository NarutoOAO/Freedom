import Modal from 'react-bootstrap/Modal';
import React, { useState, useEffect } from 'react';
import { Button } from "antd";
import { useParams } from 'react-router-dom';

function StudentDeleteSumbitAssigment(props) {
    
    const [show, setShow] = useState(false);
    const token = sessionStorage.getItem('token');
    const {courseNumber}=useParams();
    const [deleteAssigmentInfo,setDeleteAssigmentInfo]=useState('');
    const [flagForDelete,setFlagForDelete]=useState(false)
    // Open the modal
    const handleAsigmentDeleteShow = () => setShow(true);
    // Close the modal and reset input values
    const handleAsigmentDeleteClose = () => {
        setShow(false);
    };
    const fetchDeleteAssigmentInfo = async () => {
        console.log('http://127.0.0.1:5005/api/v1/assignment_submission/'+courseNumber+'/'+props.assignment_id)
        try {
          const response = await fetch('http://127.0.0.1:5005/api/v1/assignment_submission/'+courseNumber+'/'+props.assignment_id, {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
              'Authorization': token,
            },
          });
    
          if (response.status === 200) {
            const data = await response.json();
            setDeleteAssigmentInfo(data.data);
            console.log(data.data)
          } else {
            throw new Error('Failed to fetch assigment');
          }
        } catch (error) {
          console.error(error);
        }
      };
      useEffect(() => {
        if (flagForDelete) {
          fetchDeleteAssigmentInfo();
          setFlagForDelete(false); // Reset the flag to prevent further fetching
        }
         // eslint-disable-next-line
      }, [flagForDelete]);
    
    const delete_sumbit_assigment = async (props) => {
        console.log(props)
        const requestAssignmentDelete=
        {
          ass_mark_id: parseInt(props),
        };
          try {
            const response = await fetch('http://127.0.0.1:5005/api/v1/assignment_solution' , {
              method: 'DELETE',
              headers: {
                'Content-Type': 'application/json',
                'Authorization': token,
              },
              body: JSON.stringify(requestAssignmentDelete),
            });
  
            if (response.status === 200) {
              alert('Success Delete');
              fetchDeleteAssigmentInfo();
              setShow(false);
            } else {
              throw new Error('Failed to Delete');
            }
          } catch (error) {
            console.error(error);
          }
       
    };

   // ... (existing code)

return (
    <div className='createAssigment'>
      <button
        variant="primary"
        onClick={() => {
          handleAsigmentDeleteShow();
          setFlagForDelete(true); // Set the flag to trigger fetching info
        }}
        className="assigment-function-button" 
        style={{ backgroundColor: '#eceff1' }}
      >
        Delete Previous Solution
      </button>
  
      {deleteAssigmentInfo && deleteAssigmentInfo.file_url ? (
        <Modal show={show} onHide={handleAsigmentDeleteClose}>
          <Modal.Header closeButton>
            <Modal.Title>Delete Previous Solution</Modal.Title>
          </Modal.Header>
          <Modal.Body className="modalBody">
            <table>
              <thead>
                <th style={{ width: '100%' }}>The file You submit</th>
              </thead>
              <tbody>
                <a href={deleteAssigmentInfo.file_url}>{deleteAssigmentInfo.file_url}</a>
              </tbody>
            </table>
            {/* {deleteAssigmentInfo.file_url} */}
          </Modal.Body>
          <Modal.Footer>
            <Button variant="secondary" onClick={handleAsigmentDeleteClose}>
              Close
            </Button>
            <Button variant="primary" onClick={() => delete_sumbit_assigment(deleteAssigmentInfo.ass_mark_id)}>
              Delete
            </Button>
          </Modal.Footer>
        </Modal>
      )
      : (
        // Display alert if deleteAssigmentInfo is null
        <Modal show={show} onHide={handleAsigmentDeleteClose}>
          <Modal.Header closeButton>
          </Modal.Header>
          <Modal.Body>
            <p>You need to upload your assignment first.</p>
          </Modal.Body>
          <Modal.Footer>
            <Button variant="secondary" onClick={handleAsigmentDeleteClose}>
              Close
            </Button>
          </Modal.Footer>
        </Modal>
      )}
  
      {/* Display alert if deleteAssigmentInfo is null */}
      
    </div>
  )
  

}
export default StudentDeleteSumbitAssigment;