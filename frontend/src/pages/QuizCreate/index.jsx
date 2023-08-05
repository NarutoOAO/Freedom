import { Button, FormControl, Grid, InputLabel, Menu, MenuItem, Select, Stack, styled, TextField, Typography } from '@mui/material'
import React, { useEffect,useState } from 'react'
import { useAlert } from '../../hooks/useAlert';
import { Collapse } from "antd";
import QuizEdit from '../../components/QuizEdit';
import JoinQuiz from '../../components/JoinQuiz';
import { useParams} from 'react-router-dom';
import { v4 as uuidv4 } from 'uuid';
import './style.css'
const Input = styled('input')({
  display: 'none',
});
export default function QuizCreate() {
  const { alert, Alert } = useAlert();
  // const navigate = useNavigate();
  const authority = sessionStorage.getItem('authority');
  const { courseNumber } = useParams();
  const token = sessionStorage.getItem('token');
  const[quizScore, setQuizScore] = useState("")
  const [startTimeQuiz, setStartTimeQuiz] = useState('');
  const [endTimeQuiz, setEndTimeQuiz] = useState('');
  const[questions,setQuestions] = useState([{ questionId: '', answerType: '', description: '', score: '', answers: [], correctAnswerIds: [] }]
  );
  const [quizInfo, setQuizInfo] = useState([{
    name: '',
    quizScore,
    startTimeQuiz,
    endTimeQuiz
    }]);
//Collect quiz form information
  // store some info in question
  const [question, setQuestion] = useState(
    {quesionNumber:'', answerType: '', description: '', score: '', answers: '',choiceA:'',choiceB:'',choiceC:'',choiceD:'' }
  );
  
  
  const handleAddQuestion = () => {
    const formData = new FormData();
    
    const qid = String(new Date().getTime());
  
    formData.append('quiz_id', parseInt(qid));
    formData.append('question_number',parseInt(question.quesionNumber));
    formData.append('type', parseInt(question.answerType));
    formData.append('select_A', question.choiceA);
    formData.append('select_B', question.choiceB);
    formData.append('select_C', question.choiceC);
    formData.append('select_D', question.choiceD);
    formData.append('answer', question.answers);
    formData.append('description', question.description);
    formData.append('score', parseInt(question.score));
    
    fetch('http://localhost:5005/api/v1/quiz_question', {
      method: 'POST',
      headers: {
        'Authorization': token,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData)
      
    })
      .then(response => response.json())
      .then((data) => {
        if (data.status !== 200) {
          alert.warning(data.msg);
        } else {
          alert.warning('Succeed!');
        }
      })
      .catch((error) => {
        alert.warning('Failed!');
        console.log(error);
      });
      // setQuizInfo([...quizInfo, formData]);
      // console.log(setQuizInfo);
      const newQuestion = {
        questionId: parseInt(qid),
        questionNumber: question.quesionNumber,
        answerType: question.answerType,
        description: question.description,
        score: parseInt(question.score),
        answers: question.answers,
        choiceA: question.choiceA,
        choiceB: question.choiceB,
        choiceC: question.choiceC,
        choiceD: question.choiceD,
      };
    
      // Append the new question to the existing questions array
      setQuestions([...questions, newQuestion]);
    
      // Clear the form fields for adding a new question
      setQuestion({
        questionId: '',
        quesionNumber: '',
        answerType: '',
        description: '',
        score: '',
        answers: '',
        choiceA: '',
        choiceB: '',
        choiceC: '',
        choiceD: '',
      });
  };

  return (
    <div className="createQuiz" >
      <Collapse defaultActiveKey={['1']} className='panel'>
        {Alert}
        {

          authority === '1' ? (
            <div className="QuizEdit-container">
              <Button variant="contained"
                onClick={() => handleAddQuestion()}
              >Add Question</Button>
              {questions.map((item, index) => (
                <div key={index} className="quiz-panel">
                  <QuizEdit key={index} courseNumber={courseNumber} />
                </div>
              ))
              }

            </div>
          ) : (
            <div className="CompleteQuiz-container">
              <JoinQuiz courseNumber={courseNumber} />
            </div>)

        }
      </Collapse >
    </div >
  )
}