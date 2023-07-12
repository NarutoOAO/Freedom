import React, { useState } from 'react'
import logo from '../../images/search-outline.svg'
import Modal from 'react-bootstrap/Modal';

export default function PostSearch(props) {
  const [search, setSearch] = useState('');
  const [show, setShow] = useState(false);

  const courseNumber = props.courseNumber;
  const token = localStorage.getItem('token');
  const [posts, setPosts] = useState(null);

  const handleShow = async () => {
    console.log(courseNumber);
    const response = await fetch('http://127.0.0.1:5005/api/v1/post_search/'+courseNumber, {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
        'Authorization': token,
      },
      body: JSON.stringify({
        info: search
      })
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      setPosts(data.data.item);
      setShow(true);
    }
  }

  const handleClose = () => {
    setSearch('');
    setPosts(null);
    setShow(false);
  }

  const selectPost = (id) => {
    props.handlePostClickFn(id);
    handleClose();
  }

  return (
    <>
      <div style={{ display: 'flex', height: '40px' }}>
        <img src={logo} alt="search logo" style={{ width: '10%', marginLeft: '20px', marginRight: '10px' }} onClick={handleShow} />
        <input type='search' placeholder='Search' style={{ border: 'none', width: '80%' }} onChange={(event) => setSearch(event.target.value)} value={search} />
      </div>
      <Modal show={show} onHide={handleClose} size='lg'>
        <Modal.Header closeButton>
          <Modal.Title>Search List</Modal.Title>
        </Modal.Header>
        <Modal.Body className='modalBody' style={{ maxHeight: '500px', overflowY: 'auto' }}>
          <table className="table table-hover">
          <colgroup>
              <col style={{ width: '5%' }} />
              <col style={{ width: '10%' }} />
              <col style={{ width: '15%' }} />
              <col style={{ width: '70%' }} />
            </colgroup>
            <thead>
              <tr>
                <th scope="col">#</th>
                <th scope="col">Category</th>
                <th scope="col">Title</th>
                <th scope="col">Content</th>
              </tr>
            </thead>
            <tbody>
              {posts !== null &&
                posts.map((post) => (
                  <tr key={post.id} onClick={() => selectPost(post.id)}>
                    <th scope="row">{post.id}</th>
                    <td>{post.forum_name}</td>
                    <td>{post.title}</td>
                    <td>{post.content}</td>
                  </tr>
                ))}
            </tbody>
          </table>
        </Modal.Body>
      </Modal>
    </>
  )
}
