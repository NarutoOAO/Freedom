import React, { useState } from 'react'
import Modal from 'react-bootstrap/Modal';
import { Button } from 'antd';

export default function SearchMaterial(props) {
  const [searchMaterial, setSearchMaterial] = useState('');
  const [show, setShow] = useState(false);

  const courseNumber = props.courseNumber;
  const token = sessionStorage.getItem('token');
  const [postsMaterial, setPostsMaterial] = useState(null);
  const authority=sessionStorage.getItem('authority');

  const handleMaterialShow = async () => {
    console.log(authority)
    const response = await fetch('http://127.0.0.1:5005/api/v1/material/'+courseNumber, {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
        'Authorization': token,
      },
      body: JSON.stringify({
        info: searchMaterial
      })
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);

    } else {
      setPostsMaterial(data.data.item);
      setShow(true);
      document.getElementById('searchInput').value = '';
    }

  }

  const handleSearchMaterialClose = () => {
    setSearchMaterial('');
    setPostsMaterial(null);
    setShow(false);
  }

  return (
    <>
      <div style={{display: 'flex',  alignItems: 'center', width: '100%'}}>
        <input id="searchInput" type="text" name="inf" placeholder="Enter your search information" onChange={(event) => setSearchMaterial(event.target.value)}/>
        <Button onClick={handleMaterialShow}>Search</Button>
      </div>
      <Modal show={show} onHide={handleSearchMaterialClose} size='lg'>
        <Modal.Header closeButton>
          <Modal.Title>Search Material List</Modal.Title>
        </Modal.Header>
        <Modal.Body className='modalBody' style={{ maxHeight: '500px', overflowY: 'auto' }}>
          <table className="table table-hover">
          <colgroup>
              <col style={{ width: '15%' }} />
              <col style={{ width: '15%' }} />
              <col style={{ width: '70%' }} />
            </colgroup>
            <thead>
              <tr>
                <th scope="col">Week</th>
                <th scope="col">Title</th>
                <th scope="col">Link</th>
              </tr>
            </thead>
            <tbody>
            {postsMaterial !== null &&
                postsMaterial.map((post) => {
                if (authority === '0') {
                    if (post.publish === 1) {
                    return (
                        <tr key={post.id}>
                        <td>{post.file_category}</td>
                        <td>{post.file_name}</td>
                        <td>
                            <a href={post.file_url} target="_blank" rel="noopener noreferrer">
                            {post.file_url}
                            </a>
                        </td>
                        </tr>
                    );
                    } else {
                    return null; 
                    }
                } else if (authority === '1') {
                    return (
                    <tr key={post.id}>
                        <td>{post.file_category}</td>
                        <td>{post.file_name}</td>
                        <td>
                        <a href={post.file_url} target="_blank" rel="noopener noreferrer">
                            {post.file_url}
                        </a>
                        </td>
                    </tr>
                    );
                } else {
                    return null; 
                }
                })}
            </tbody>

          </table>
        </Modal.Body>
      </Modal>
    </>
  )
}
