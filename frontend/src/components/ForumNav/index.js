import React, { useState, useEffect } from 'react';
import Button from 'react-bootstrap/Button';
import ExitToAppIcon from '@mui/icons-material/ExitToApp';
import { useNavigate } from 'react-router-dom';
import CreateCategories from '../../components/CreateCategories';
import { useParams } from 'react-router-dom';
import PostForum from '../PostForum';
import './style.css';

export default function ForumNav(props) {
  const navigate = useNavigate();
  const [categories, setCategories] = useState([]);
  const token = sessionStorage.getItem('token');
  const { courseNumber } = useParams();
  const [flag, setFlag] = useState(0);
  const authority=sessionStorage.getItem('authority');

  useEffect(() => {
    fetchCategories();
    setFlag(0);
    // eslint-disable-next-line
  },[courseNumber,flag]);

  const goBack = () => {
    navigate(-1);
  };

  const fetchCategories = async () => {
    const response = await fetch(`http://127.0.0.1:5005/api/v1/forum/${courseNumber}`, {
      method: 'GET',
      headers: {
        'Content-type': 'application/json',
        Authorization: token,
      },
    });
    const data = await response.json();
    if (data.status !== 200) {
      alert(data.msg);
    } else {
      setCategories(data.data);
    }
  };

  const handleCategoryClick = (categoryId) => {
    props.setSelectedCateFn(categoryId);
  };

  return (
    <div className="navigation_teacher">
      <div className='EDNavTitle'>
        <div>
          Discussion Forum
        </div>
        <div>
          {courseNumber}
        </div>
      </div>
      <ul>
        <li>
          <div>
            <Button onClick={goBack} style={{ width: '80%', height: '40px', fontSize: '10pt' }}>
              <ExitToAppIcon /> Back
            </Button>
          </div>
          {authority !== '0' && (
          <div>
            <CreateCategories courseNumber={courseNumber} setFlag={setFlag}/>
          </div>
          )}
          <PostForum courseNumber={courseNumber} categories={categories}/>
          <div style={{ fontSize: 'small', marginTop: '20px', textAlign: 'left', marginLeft:'20px' }}>
            CATEGORIES
          </div>
        </li>
        {categories !== null &&
          categories.map((category) => (
            <div
              key={category.ID}
              style={{ fontSize: '10pt', textAlign: 'left', marginLeft: '30px', marginTop: '10px', width:'70%'}}
              onClick={() => handleCategoryClick(category.ID)} className='NavCategory'
            >
              {category.ForumName}
            </div>
          ))}
      </ul>
    </div>
  );
}
