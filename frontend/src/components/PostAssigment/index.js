import React, { useState } from 'react';
import { Button } from "antd";
import Modal from 'react-bootstrap/Modal';

function PostAssigment(props) {
  const [show, setShow] = useState(false);
  const [inputAssigmentTitle, setAssigmentInputTitle] = useState("");
  const [maxScore,setMaxScore]= useState("");
  const courseNumber = props.courseNumber;
  // Open the modal
  const handleShow = () => setShow(true);
  // Close the modal and reset input values
  const handleClose = () => {
    setShow(false);
    setAssigmentInputTitle('');
    setMaxScore('');
  };


  // Handle assignment title input change
  const handleInputTitleChange = (e) => {
    setAssigmentInputTitle(e.target.value);
  };

  // Handle maximum score input change
  const handleMaxScoreChange = (e) => {
    setMaxScore(e.target.value);
  };

  // Handle assignment upload
  function handleAssigmentUpload(props) {
    const fileInput = document.getElementById('fileInput');
    const file = fileInput.files[0];
    const formData = new FormData();
    const token = sessionStorage.getItem('token');

    formData.append('course_number', parseInt(courseNumber));
    formData.append('file_name', inputAssigmentTitle);
    formData.append('file', file);
    formData.append('max_score',parseInt(maxScore));
    const maxScoreValue = parseInt(maxScore);

    // Validate inputAssigmentTitle, file, and maxScore
    if (!inputAssigmentTitle || !file || !maxScore) {
      alert('Please fill in all the required fields.');
      return;
    }
    // Validate max score
    if (isNaN(maxScoreValue) || maxScoreValue < 0 || maxScoreValue > 100) {
      alert('Max score must be between 0 and 100.');
      return;
    }

    // Validate file type
    if (file.type !== 'application/pdf') {
      alert('Only PDF files are allowed.');
      return;
    }
   
    // Send the assignment data to the server
    fetch('http://127.0.0.1:5005/api/v1/assignment', {
      method: 'POST',
      headers: {
        'Authorization': token,
      },
      body:formData,
    })
      .then(response => response.json())
      .then((data) => {
        if (data.status !== 200) {
          alert(data.msg);
        } else {
          alert('Succeed!');
          props.setAssigmentFlag(1);
          handleClose();
        }
      }) 
      .catch((error) => {
        alert("Failed!")
      });
      
  }
  

  return (
    <div className='createAssigment'>
      <Button
        variant="primary"
        onClick={handleShow}
        className="materialBtn"
        size="large"
        style={{ borderRadius: '8px', boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)' }}
      >
        Create New Assigment
      </Button>

      <Modal show={show} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Create New Assigment</Modal.Title>
        </Modal.Header>
        <Modal.Body className="modalBody">
          <form encType="multipart/form-data">
            <div>
              <label htmlFor="inputTitle">Title:</label> <br />
              <input type="text" id="inputTitle" value={inputAssigmentTitle} onChange={handleInputTitleChange} style={{width:'100%', marginBottom:'10px'}}/>
            </div>
            <div>
              <label htmlFor="inputScore">Score:</label> <br />
              <input type="text" id="inputScore" value={maxScore} onChange={handleMaxScoreChange} style={{width:'100%', marginBottom:'10px'}}/>
            </div>
            <div>
              <label htmlFor="fileInput">File:</label><br />
              <input type="file" id="fileInput" />
              <div style={{ color: 'gray', fontStyle: 'italic' ,marginTop:'10px'}}>Only PDF files are allowed</div>
            </div>
            
          </form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button variant="primary" onClick={() => handleAssigmentUpload(props)}>
              Upload
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
}

export default PostAssigment;
