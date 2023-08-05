import React, { useState} from 'react'
import ForumOverview from '../../components/ForumOverview'
import ForumNav from '../../components/ForumNav'
import { useParams } from 'react-router-dom';
import ForumContent from '../../components/ForumContent'
import './style.css'
// define the discussion forum
export default function DiscussionForum() {
  const { courseNumber } = useParams();
  const [selectedCate,setSelectedCate] = useState('-1');
  const [selectedPost, setSelectedPost] = useState('');
  const [flag, setFlag] = useState(0);
  return (
    <div className="discussion-container">
      <ForumNav courseNumber={courseNumber} setSelectedCateFn={setSelectedCate} setFlag={setFlag}/>
      <ForumOverview courseNumber={courseNumber} selectedCate={selectedCate}  setSelectedPostFn={setSelectedPost}/>
      <ForumContent courseNumber={courseNumber} selectedPost={selectedPost}/>
    </div>
  )
}