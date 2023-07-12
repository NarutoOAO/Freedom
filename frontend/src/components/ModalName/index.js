import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';

function ModalName(props) {
  const [show, setShow] = useState(false);
  const token = localStorage.getItem('token');
  const name = localStorage.getItem('name');
  const [newName,setNewName] = useState('');
  const handleClose = () => {
    setShow(false);
    setNewName('');
  }
  const handleShow = () => setShow(true);
  const modify = async () => {
    const response = await fetch('http://localhost:5005/api/v1/user', {
      method: 'PUT',
      headers: {
        'Content-type': 'application/json',
        'Authorization': token,
      },
      body: JSON.stringify({
        nickname:newName,
      })
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      props.setNameFn(data.data.nickname);
      localStorage.setItem('name', data.data.nickname);
      alert("Succeed!");
      window.location.reload();
      handleClose();
    }
  }
  return (
    <>
      <Button variant="primary" onClick={handleShow}>
      Modify name
      </Button>

      <Modal show={show} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Modify name</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Label>Current name</Form.Label>
              <div>
                {name}
              </div>
            </Form.Group>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Label>New name</Form.Label>
              <Form.Control
                type="text"
                autoFocus
                onChange={(event) => setNewName(event.target.value)} value={newName}
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

export default ModalName;