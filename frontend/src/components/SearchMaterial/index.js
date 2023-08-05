import React, { useState } from 'react'
import Modal from 'react-bootstrap/Modal';
import { Button } from 'antd';
// Define a component to search material
export default function SearchMaterial(props) {
  // Use to store search material data
  const [searchMaterial, setSearchMaterial] = useState('');
  const [show, setShow] = useState(false);
  const courseNumber = props.courseNumber;
  // Get user token from session storage
  const token = sessionStorage.getItem('token');
  const [postsMaterial, setPostsMaterial] = useState(null);
  // Get user authority from session storage
  const authority=sessionStorage.getItem('authority');
  // Open the modal and fetching search results
  const handleMaterialShow = async () => {
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
  
  // Close the modal and reset input values
  const handleSearchMaterialClose = () => {
    setSearchMaterial('');
    setPostsMaterial(null);
    setShow(false);
  }

  return (
    <>
      <div style={{display: 'flex',  alignItems: 'center', width: '100%'}}>
        {/* Search input field */}
        <input id="searchInput" type="text" name="inf" placeholder="Enter your search information" onChange={(event) => setSearchMaterial(event.target.value)}/>
        {/* Search Button */}
        <Button onClick={handleMaterialShow}>Search</Button>
      </div>
      {/* Modal component to display search results */}
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
            {/* Search results based on user authority */}
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
