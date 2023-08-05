import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';
// define a component to post the comment
function PostComment(props) {
  const [show, setShow] = useState(false);
  const token = sessionStorage.getItem('token');
  // const name = sessionStorage.getItem('name');
  const [comment,setComment] = useState('');
  const handleClose = () => {
    setShow(false);
    setComment('');
  }
  const handleShow = () => setShow(true);
  // Handle comment input change
  const newComment = async () => {
    const response = await fetch('http://127.0.0.1:5005/api/v1/comment', {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
        'Authorization': token,
      },
      body: JSON.stringify({
        post_id:parseInt(props.post_id),
        content:comment,
      })
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      // console.log(data.data);
      alert("Succeed!");
      props.setFlag(1);
      handleClose();
    }
  }
  return (
    <>
      <span onClick={handleShow}>
      Comment ...
      </span>

      <Modal show={show} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Create a new comment</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Label>Comment</Form.Label>
              <Form.Control
                type="text"
                autoFocus
                onChange={(event) => setComment(event.target.value)} value={comment}
              />
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button variant="primary" onClick={newComment}>
            Submit
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
}

export default PostComment;