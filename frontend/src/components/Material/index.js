import React, { useState } from 'react';
import { Button } from "antd";
import Modal from 'react-bootstrap/Modal';

function PostMaterial(props) {
  const [show, setShow] = useState(false);
  const [selectedWeek, setSelectedWeek] = useState("");
  const [selectedFileType, setSelectedFileType] = useState("");
  const [inputTitle, setInputTitle] = useState("");
  const courseNumber = props.courseNumber;

  const handleShow = () => setShow(true);
  const handleClose = () => {
    setShow(false);
    setSelectedWeek('');
    setSelectedFileType('');
    setInputTitle('');
  };

  const handleWeekChange = (e) => {
    setSelectedWeek(e.target.value);
  };

  const handleFileTypeChange = (e) => {
    setSelectedFileType(e.target.value);
  };

  const handleInputTitleChange = (e) => {
    setInputTitle(e.target.value);
  };
  
  function handleUpload() {
    const fileInput = document.getElementById('fileInput');
    const file = fileInput.files[0];
    const formData = new FormData();
    const token = localStorage.getItem('token');

    formData.append('course_number', parseInt(courseNumber));
    formData.append('file_name', inputTitle);
    formData.append('file', file);
    formData.append('file_category', selectedWeek);
    formData.append('type', parseInt(selectedFileType));
    formData.append('publish',parseInt('2'));
   
    fetch('http://127.0.0.1:5005/api/v1/material', {
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
          handleClose();
        }
      }) 
      .catch((error) => {
        alert("Failed!")
      });
      
  }
  

  return (
    <div className='createMaterial'>
      <Button
        variant="primary"
        onClick={handleShow}
        className="materialBtn"
        size="large"
        style={{ borderRadius: '8px', boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)' }}
      >
        Create New Material
      </Button>

      <Modal show={show} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Create New Material</Modal.Title>
        </Modal.Header>
        <Modal.Body className="modalBody">
          <form encType="multipart/form-data">
            <div>
              <label htmlFor="inputTitle">Title:</label> <br />
              <input type="text" id="inputTitle" value={inputTitle} onChange={handleInputTitleChange} style={{width:'100%', marginBottom:'10px'}}/>
            </div>
            <div>
              <label htmlFor="weekSelect">Select Week:</label><br />
              <select id="weekSelect" value={selectedWeek} onChange={handleWeekChange} style={{width:'100%', fontSize:'17pt', marginBottom:'10px'}}>
                <option value="">Select Week</option>
                <option value="Week 1">Week 1</option>
                <option value="Week 2">Week 2</option>
                <option value="Week 3">Week 3</option>
                <option value="Week 4">Week 4</option>
                <option value="Week 5">Week 5</option>
                <option value="Week 6">Week 6</option>
                <option value="Week 7">Week 7</option>
                <option value="Week 8">Week 8</option>
                <option value="Week 9">Week 9</option>
                <option value="Week 10">Week 10</option>
              </select>
            </div>
            <div>
              <label htmlFor="fileTypeSelect">Select File Type:</label><br />
              <select id="fileTypeSelect" value={selectedFileType} onChange={handleFileTypeChange} style={{width:'100%', fontSize:'17pt', marginBottom:'10px', fontWeight:'normal'}}>
                <option value="">Select File Type</option>
                <option value="0">PDF</option>
                <option value="1">PPT</option>
              </select>
            </div>
            <div>
              <label htmlFor="fileInput">File:</label><br />
              <input type="file" id="fileInput" />
            </div>
          </form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button variant="primary" onClick={handleUpload}>
            Upload
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
}

export default PostMaterial;
