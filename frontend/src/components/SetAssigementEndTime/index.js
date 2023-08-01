import { Button, Modal } from 'react-bootstrap';
import React, { useState } from 'react';

function SetAssigementEndTime(props) {
  const token = sessionStorage.getItem('token');
  const [showModal, setShowModal] = useState(false);
  const [selectedAssignmentId, setSelectedAssignmentId] = useState(null);
  const [endTimeInput, setEndTimeInput] = useState('');

  const handleShowModal = (assignmentId) => {
    setSelectedAssignmentId(assignmentId);
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setSelectedAssignmentId(null);
    setEndTimeInput('');
  };

 // Function to format the selected end time  
  const formatDateTime = (dateTime) => {
    const date = new Date(dateTime);
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const seconds = String(date.getSeconds()).padStart(2, '0');
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
  };

  const setEndTime = async (props) => {
    // Format the selected end time
    const formattedEndTime = formatDateTime(endTimeInput);

    // Check if the end time is earlier than the current time
    const currentDateTime = new Date();
    if (new Date(formattedEndTime) < currentDateTime) {
      alert('End time cannot be earlier than the current time.');
      return;
    }
    const requestPublish = {
      assignment_id: props.assignment_id,
      end_time: formattedEndTime,
    };

    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/assignment', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
        body: JSON.stringify(requestPublish),
      });

      if (response.status === 200) {
        props.props.setAssigmentFlag(1);
        handleCloseModal();
      } else {
        throw new Error('Failed to Publish');
      }
    } catch (error) {
      console.error(error);
    }
  };

  const handleDateTimeInputChange = (e) => {
    setEndTimeInput(e.target.value);
  };

  return (
    <div>
      {/* Button to open the modal */}
      <button
        className="assigment-function-button"
        style={{ backgroundColor: '#cfd8dc' }}
        onClick={() => handleShowModal(props.assignment_id)}
      >
        Set End time
      </button>

      {/* Modal component */}
      <Modal show={showModal} onHide={handleCloseModal}>
        <Modal.Header closeButton>
          <Modal.Title>Set End Time</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <label htmlFor="endTimeInput" >End Time:</label>
          <input
            type="datetime-local"
            id="myDateTimePicker"
            value={endTimeInput}
            onChange={handleDateTimeInputChange}
            style={{ marginLeft: '20px' }}
          />
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleCloseModal}>
            Cancel
          </Button>
          <Button variant="primary" onClick={() => setEndTime(props)}>
            Save
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
}

export default SetAssigementEndTime;
