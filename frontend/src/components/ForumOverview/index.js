import React, { useState, useEffect } from 'react';
import ListGroup from 'react-bootstrap/ListGroup';
import SearchPosts from '../SearchPosts';
import './style.css'
export default function ForumOverview(props) {
  const [posts, setPosts] = useState([]);
  const token = sessionStorage.getItem('token');
  const selectedCate = props.selectedCate;
  const courseNumber = props.courseNumber;

  useEffect(() => {
    if(selectedCate==='-1'){
      fetchPosts();
    }else{
      fetchPostsByPostId();
    }
    // eslint-disable-next-line
  },[selectedCate]);

  const fetchPosts = async () => {
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/posts/'+courseNumber, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
      });

      if (response.status === 200) {
        const data = await response.json();
        if(data.status===200){
          // console.log(data);
          setPosts(data.data);
        }
        else{
          alert(data.msg);
        }
      } else {
        throw new Error('Failed to fetch posts');
      }
    } catch (error) {
      console.error(error);
    }
  };

  const fetchPostsByPostId = async () => {
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/post/'+selectedCate, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
      });

      if (response.status === 200) {
        const data = await response.json();
        if(data.status===200){
          // console.log(data);
          setPosts(data.data);
        }
        else{
          alert(data.msg);
        }
      } else {
        throw new Error('Failed to fetch posts');
      }
    } catch (error) {
      console.error(error);
    }
  };

  const handlePostClick = (postId) => {
    // console.log(postId);
    props.setSelectedPostFn(postId);
  };

  return (
    <div className='forumOverList'>
      <SearchPosts courseNumber = {courseNumber} handlePostClickFn = {handlePostClick}/>
      <div style={{backgroundColor:'#f2f2f2', textAlign:'center', fontSize:'small', padding:'10px'}}>Posts</div>
      {posts!==null ? (
        <ListGroup style={{height:'100%'}}>
        {posts.map((post) => (
            <ListGroup.Item key={post.id} value={post.id} onClick={() => handlePostClick(post.id)} >
              <div style={{ margin:'5px'}}>
              {post.title}
              </div>
            </ListGroup.Item>
          ))}
        </ListGroup>
      ) : (
        <div></div>
      )}
    </div>
  );
}
