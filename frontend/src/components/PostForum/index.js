import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';

function PostForum(props) {
  const [show, setShow] = useState(false);
  const token = sessionStorage.getItem('token');
  const [title,setTitle] = useState('');
  const [content,setContent] = useState('');
  const categories = props.categories;
  const [category, setCategory] = useState(1);
  const handleClose = () => {
    setShow(false);
    setCategory(1);
    setContent('');
    setTitle('');
  }
  const handleShow = () => setShow(true);
  const createBtn = async () => {
    const response = await fetch('http://127.0.0.1:5005/api/v1/post', {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
        'Authorization': token,
      },
      body: JSON.stringify({
        forum_id:parseInt(category),
        title:title,
        content:content,
      })
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      alert('success!');
      handleClose();
    }
  }

  return (
    <>
      <Button  style={{ width: '80%', height: '40px', fontSize:'10pt' }} onClick={handleShow} className='createCatBtn'>
      New Thread
      </Button>
      <Modal show={show} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Create a new question</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
          <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Label>Title</Form.Label>
              <Form.Control
                type="text"
                autoFocus
                onChange={(event) => setTitle(event.target.value)} value={title}
              />
            </Form.Group>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Label>Content</Form.Label>
              <Form.Control
                type="text"
                autoFocus
                onChange={(event) => setContent(event.target.value)} value={content}
              />
            </Form.Group>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Label>Category</Form.Label>
              <Form.Select
                value={category}
                onChange={(event) => setCategory(event.target.value)}
              >
                {categories !== null ? (
                  categories.map((categoryItem) => (
                    <option key={categoryItem.ID} value={categoryItem.ID}>
                      {categoryItem.ForumName}
                    </option>
                  ))
                ) : (
                  <option disabled>You need to create categories first!</option>
                )}
              </Form.Select>
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button variant="primary" onClick={createBtn}>
            Post
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
}

export default PostForum;