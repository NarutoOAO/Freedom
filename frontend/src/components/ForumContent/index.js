import React, { useState, useEffect } from 'react'
import logo from '../../images/person-circle-outline.svg'
import heart from '../../images/heart-outline.svg'
import PostComment from '../PostComment'
import './style.css'
export default function ForumContent(props) {
  const selectedPost = props.selectedPost;
  const token = sessionStorage.getItem('token');
  const [post, setPost] = useState('');
  const [contents,setContents] = useState('');
  const [flag, setFlag] = useState(0);

  useEffect(() => {
    if(selectedPost!==''){
      fetchPost();
      fetchContents();
    }
    setFlag(0);
  },[selectedPost,flag]);

  const fetchPost = async () => {
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/post_information/'+selectedPost, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
      });

      if (response.status === 200) {
        const data = await response.json();
        // console.log(data.data);
        setPost(data.data);
      } else {
        throw new Error('Failed to fetch post');
      }
    } catch (error) {
      console.error(error);
    }
  };

  const fetchContents = async () => {
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/comment/'+selectedPost, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
      });

      if (response.status === 200) {
        const data = await response.json();
        // console.log(data.data);
        setContents(data.data);
      } else {
        throw new Error('Failed to fetch post');
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <>
    {selectedPost !== '' && (<div className='forumContent'>
      <div className='contentTitle'>
        <h2> {post.title}</h2>
        <label>#{post.id}</label>
      </div>
      <div className='contentAuthor'>
        <div>
        <img src={logo} alt="logo"/> 
        </div>
        <div>
          <div>
            {post.author_name}
          </div>
          <div>
            {post.forum_name}
          </div>
        </div>
      </div>
      <div className='contents'>
        <div>
        <img src={heart} alt="heart"/> 
        </div>
        <p>
          {post.content}
        </p>
      </div>
      <PostComment post_id={selectedPost} setFlag={setFlag}/>
      <h2 style={{marginTop:'40px', marginBottom:'20px'}}>Answer</h2>
      <ul>
      {Array.isArray(contents) &&
          contents.map((content) => (
          <li key={content.id}><img src={logo} alt="logo"/> <span>{content.author_name}</span>{content.authorization?(<label>Tutor</label>):<></>}<p>{content.content}</p></li>
          ))}
      </ul>
    </div>)}
    </>
  )
}
