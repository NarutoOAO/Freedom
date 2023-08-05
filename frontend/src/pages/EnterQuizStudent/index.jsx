import React, { useEffect, useState } from 'react';
import { Card, List } from 'antd';
// import { useAlert } from '../../hooks/useAlert';
import ModalJoinQuiz from '../../components/ModalJoinQuiz';
import { Button } from 'react-bootstrap';
import { useParams } from 'react-router-dom';

export default function EnterQuizStudent(props) {
  const [quizInfo, setQuizInfo] = useState([]);
  // const { alert, Alert } = useAlert();
  const token = sessionStorage.getItem('token');
  const { courseNumber } = useParams();
  const [showModal, setShowModal] = useState(false); // State to control modal visibility
  const [selectedQuizId, setSelectedQuizId] = useState(null); // State to store the selected quiz ID

  useEffect(() => {
    showInfo(); // Call the function when the component mounts
  }, []);


  const showInfo = async () => {
    try {
      const response = await fetch('http://127.0.0.1:5005/api/v1/quiz/' + courseNumber, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token,
        },
      });

      if (response.status === 200) {
        const data = await response.json();
        if (data.status === 200) {
          setQuizInfo(data.data);
        } else {
          alert(data.msg);
        }
      } else {
        throw new Error('Failed to fetch posts');
      }
    } catch (error) {
      alert(error);
    };
  }

  const handleEnterQuiz = (quizId,start_time,end_time) => {

    setSelectedQuizId(quizId);
    const currentTime = new Date();
    if (currentTime > new Date(end_time)) {
      alert('Sorry, the quiz has already ended. You cannot participate anymore.');
      return;
    } else if (new Date(start_time) < currentTime && currentTime<new Date(end_time)) {
      setShowModal(true);
    } else {
      alert('The quiz not start yet.');
      return;
      
    }
  }

  const handleCloseModal = () => {
    setShowModal(false);
  };

  return (
    <div style={{ height: '600px', overflowY: 'auto' }}>
      {/* {Alert} */}
      {quizInfo && quizInfo.length > 0 ? (
        <Card className="quiz-info-card" >
          <List >
            {/* using map to iterate all information of the quiz such as quiz name, score and so on */}
            {quizInfo.map((quiz, index) => (
              <React.Fragment key={index}>
                <List.Item>Quiz Name: {quiz.quiz_name}</List.Item>
                <List.Item>Score: {quiz.max_score}</List.Item>
                <List.Item>Start Time: {quiz.start_time}</List.Item>
                <List.Item>End Time: {quiz.end_time}</List.Item>
                <Button variant="primary" color="blue" onClick={() => handleEnterQuiz(quiz.quiz_id,quiz.start_time,quiz.end_time)}>Enter</Button>
              </React.Fragment>
            ))}
          </List>
        </Card>
      ) : (
        <p>No quiz found</p>
      )}
      <div>
        {selectedQuizId && (
          <ModalJoinQuiz quiz_id={selectedQuizId} course_number={courseNumber} showModal={showModal} onClose={handleCloseModal} />
        )}
      </div>
    </div>
  )
}