import Modal from 'react-bootstrap/Modal';
import React, { useState } from 'react';
import { Button } from "antd";
import { useParams } from 'react-router-dom';
import AllocateGroup from '../AllocateGroup';
// define a component to mark the assignment
function MarkAssigment(props) {
  const [show, setShow] = useState(false);
  const [mark, setMark] = useState({});
  const [content, setContent] = useState({});
  const [assigmentInfroamtion, setAssigmentInfroamtion] = useState([]);
  const token = sessionStorage.getItem('token');
  const { courseNumber } = useParams();

  // Open the modal
  const handleAsigmentUploadShow = () => setShow(true);

  // Close the modal and reset input values
  const handleAsigmentUploaClose = () => {
    setMark({});
    setContent({});
    setShow(false);
  };

  // Handle input change for mark field
  const handleMarkChange = (event, assMarkId) => {
    setMark(prevState => ({
      ...prevState,
      [assMarkId]: event.target.value
    }));
  };

  // Handle input change for content field
  const handleContentChange = (event, assMarkId) => {
    setContent(prevState => ({
      ...prevState,
      [assMarkId]: event.target.value
    }));
  };

  // Fetch assignment data when marking is initiated
  const mark_assigment = async () => {
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/assignment_solution/' + courseNumber + '/' + parseInt(props.assignment_id), {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
      });

      if (response.status === 200) {
        const data = await response.json();
        if (data.data === null) {
          alert('No one posted the assignment');
          handleAsigmentUploaClose();
          return;
        } else {
          setAssigmentInfroamtion(data.data);
          handleAsigmentUploadShow();
        }
      } else {
        throw new Error('Failed to fetch assignment');
      }
    } catch (error) {
      console.error(error);
    }
  };

  // Submit marked assignment  
  const mark_function = async (assMarkId, max_score) => {
    const markValue = parseInt(mark[assMarkId]);
    
    // Check if markValue is a valid number and within the valid range
    if (isNaN(markValue) || markValue <= 0 || markValue > max_score) {
      alert('Please enter a valid mark');
      return;
    }
    
    const markInfotmaion = {
      ass_mark_id: parseInt(assMarkId),
      mark: markValue,
      content: content[assMarkId],
    };
  
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/assignment_grade', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
        body: JSON.stringify(markInfotmaion),
      });
  
      if (response.status === 200) {
        alert('Scores have been marked successfully.');
        setMark(prevState => ({
          ...prevState,
          [assMarkId]: ''
        }));
        setContent(prevState => ({
          ...prevState,
          [assMarkId]: ''
        }));
      } else {
        throw new Error('Failed to Publish');
      }
    } catch (error) {
      console.error(error);
    }
  };

  // Function to render either the AllocateGroup component or group_name
  const renderAddorGroup=(ass)=>{
    if (ass.group_id===0){
      return <AllocateGroup courseNumber={courseNumber} id={ass.ass_mark_id} mark_assigmentFn={mark_assigment}/>
    }else{
      return <span>{ass.group_name}</span>
    }
  }

  return (
    <div className='createAssigment'>
      <button className="assigment-function-button" style={{ backgroundColor: '#eceff1' }} onClick={() => mark_assigment(props)}>Mark</button>
      <Modal size='xl' show={show} onHide={handleAsigmentUploaClose}>
        <Modal.Header closeButton>
          <Modal.Title>Mark Assigment</Modal.Title>
        </Modal.Header>
        <Modal.Body className='modalBody' style={{ overflowY: 'auto', overflowX: 'auto' }}>
          <table className="table table-hover">
            <colgroup>
              <col style={{ width: '5%' }} />
              <col style={{ width: '30%' }} />
              <col style={{ width: '5%' }} />
              <col style={{ width: '30%' }} />
              <col style={{ width: '15%' }} />
              <col style={{ width: '15%' }} />
            </colgroup>
            <thead>
              <tr>
                <th scope="col">Student Name</th>
                <th scope="col">Link</th>
                <th scope="col">Mark</th>
                <th scope="col">Feedback</th>
                <th scope="col">Button</th>
                <th scope="col">Group</th>
              </tr>
            </thead>
            <tbody>
              {assigmentInfroamtion !== null &&
                assigmentInfroamtion.map((assigmentInfroamtion) => {
                  const assMarkId = assigmentInfroamtion.ass_mark_id;
                  return (
                    <tr key={assMarkId}>
                      {/* Map through assignment information and display them in a table */}
                      <td>{assigmentInfroamtion.file_url.substring(assigmentInfroamtion.file_url.lastIndexOf("/") + 1, assigmentInfroamtion.file_url.lastIndexOf("."))}</td>
                      <td>
                        <a href={assigmentInfroamtion.file_url} target="_blank" rel="noopener noreferrer">
                          {assigmentInfroamtion.file_url}
                        </a>
                      </td>
                      <td>
                        <input type="text" placeholder='Mark' value={mark[assMarkId] || ''} onChange={(event) => handleMarkChange(event, assMarkId)}/>
                      </td>
                      <td>
                        <input type="text" placeholder='Feedback' value={content[assMarkId] || ''} onChange={(event) => handleContentChange(event, assMarkId)} />
                      </td>
                      <td>
                        <Button style={{ backgroundColor: "#87CEEB" }} onClick={() => mark_function(assMarkId,assigmentInfroamtion.max_score)}>Mark</Button>
                      </td>
                      <td>
                        {renderAddorGroup(assigmentInfroamtion)}
                      </td>
                    </tr>
                  );
                })}
            </tbody>
          </table>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleAsigmentUploaClose}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
}

export default MarkAssigment;
