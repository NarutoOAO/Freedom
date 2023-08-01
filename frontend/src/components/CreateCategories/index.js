
import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';
import './style.css'
function CreateCategories(props) {
  const [show, setShow] = useState(false);
  const token = sessionStorage.getItem('token');
  const [category,setCategory] = useState('');
  const courseNumber = props.courseNumber;

  const handleClose = () => {
    setShow(false);
    setCategory('');
  }
  const handleShow = () => setShow(true);
  const createBtn = async () => {
    const response = await fetch('http://127.0.0.1:5005/api/v1/forum', {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
        'Authorization': token,
      },
      body: JSON.stringify({
        forum_name:category,
        course_number: parseInt(courseNumber)
      })
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      alert('success!');
      props.setFlag(1);
      handleClose();
    }
  }
  return (
    <>
      <Button  style={{ width: '80%', height: '40px', fontSize:'10pt'}} onClick={handleShow} className='createCatBtn'>
      New Category
      </Button>

      <Modal show={show} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Create a new category</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
          <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Label>Category name</Form.Label>
              <Form.Control
                type="text"
                autoFocus
                onChange={(event) => setCategory(event.target.value)} value={category}
              />
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button variant="primary" onClick={createBtn}>
            Create
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
}

export default CreateCategories;