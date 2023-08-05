import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';
// define a component to create a new group
function CreateGroup(props) {
  const [show, setShow] = useState(false);
  const token = sessionStorage.getItem('token');
  const courseNumber = props.courseNumber;
  const name = sessionStorage.getItem('name');
  const user_id = sessionStorage.getItem('user_id');
  const [group, setGroup] = useState('');
  // handle close and show
  const handleClose = () => {
    setShow(false);
    setGroup('');
  }
  const handleShow = () => setShow(true);
  // create a new group
  const modify = async () => {
    const response = await fetch('http://localhost:5005/api/v1/group', {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
        'Authorization': token,
      },
      body: JSON.stringify({
        course_number: parseInt(courseNumber),
        group_name: group,
        teacher_id: parseInt(user_id),
        teacher_name: name,
      })
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      alert("Succeed!");
      props.getGroupsFn();
      handleClose();
    }
  }
  return (
    <>
      <Button variant="primary" onClick={handleShow} style={{marginBottom:"20px"}}>
          Create Group
      </Button>

      <Modal show={show} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Create Group</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Label>New group</Form.Label>
              <Form.Control
                type="text"
                autoFocus
                onChange={(event) => setGroup(event.target.value)} value={group}
              />
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button variant="primary" onClick={modify}>
            Submit
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
}

export default CreateGroup;