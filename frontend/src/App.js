import { Routes, Route, useNavigate, useLocation } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import NavScrollExample from './components/NavBar'
import Layout from './pages/Layout';
import Login from './pages/Login';
import Register from './pages/Register';
import React, { useEffect, useState } from 'react';
import Dashboard1 from './pages/student/dashboard';
import Dashboard2 from './pages/teacher/dashboard';
import Profile from './pages/Profile'
import TeacherCourse from './pages/teacher/course'
import StudentCourse from './pages/student/course'
import CreateCourse from './pages/CreateCourse'
import EnrollCourse from './pages/EnrollCourse'
import DiscussionForum from './pages/DiscussionForum'
import ForumPosts from './components/ForumOverview'
import QuizCreate from './pages/QuizCreate'
import './App.css';

function App() {
  const [token, setToken] = useState(sessionStorage.getItem('token'));
  const navigate = useNavigate();
  const { pathname } = useLocation();

  useEffect(() => {
    if (token !== null && token !== 'null' && (pathname === '/login' || pathname === '/register')) {
      document.body.style.background='white';
      const authority = parseInt(sessionStorage.getItem('authority'));
      if(authority === 0){
        navigate('/student/dashboard');
      }
      else{
        navigate('/teacher/dashboard');
      }
    }
  }, [token, navigate, pathname]);

  return (
    <div className="App">
      {pathname!=='/' &&(<NavScrollExample setTokenFn={setToken}/>)}
      <Routes>
        <Route path="/" element={<Layout />} />
        <Route path="/login" element={<Login  setTokenFn={setToken}/>}/>
        <Route path="/register" element={<Register  setTokenFn={setToken}/>}/>
        <Route path="/student/dashboard" element={<Dashboard1  setTokenFn={setToken}/>} />
        <Route path="/teacher/dashboard" element={<Dashboard2  setTokenFn={setToken}/>} />
        <Route path="/profile" element={<Profile/>} />
        <Route path='/teacher/course/:courseNumber' element={<TeacherCourse/>} />
        <Route path='/teacher/course' element={<QuizCreate/>} />
        <Route path='/student/course' element={<QuizCreate/>} />
        <Route path='/student/course/:courseNumber' element={<StudentCourse/>} />
        <Route path='/create_course' element={<CreateCourse/>} />
        <Route path='/enroll_course' element={<EnrollCourse/>} />
        <Route path='/discussion_forum/:courseNumber' element={<DiscussionForum/>} />
        <Route path='/forum_overview/:courseNumber' element={<ForumPosts/>} />
      </Routes>
    </div>
  );
}

export default App;
