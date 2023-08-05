import React, { useState, useEffect } from 'react';
import { Input, Form } from "antd";
import { useParams } from 'react-router-dom';
import QuizEdit from '../../components/QuizEdit';
//import EnterQuizStudent from '../EnterQuizStudent';

export default function EnterQuiz() {
  const { courseNumber } = useParams();
  //const MyContext = React.createContext();
  const layout = {
    labelCol: { span: 8 },
    wrapperCol: { span: 16 },
  };

  const [form] = Form.useForm();
  // const[inputTitle, setInputTitle] = useState("");
  const [quizName, setQuizName] = useState("");
  const [quizScore, setQuizScore] = useState("")
  const [startTimeQuiz, setStartTimeQuiz] = useState('');
  const [endTimeQuiz, setEndTimeQuiz] = useState('');
  const [flagForClean, setFlagForClean] = useState(false);

  // the xxxchange functions are using for storing input values
  const handleInputTitleChange = (e) => {
    setQuizName(e.target.value);
  }
  const handleInputScoreChange = (e) => {
    setQuizScore(e.target.value);
  }
  const handleStartQuizChange = (e) => {
    setStartTimeQuiz(e.target.value);
  };
  const handleEndQuizChange = (e) => {
    setEndTimeQuiz(e.target.value);
  };
  const handleClean = () => {
    // clean all these info
    setQuizName("");
    setQuizScore("");
    setStartTimeQuiz("");
    setEndTimeQuiz("");

  };
  useEffect(() => {
    if (flagForClean) {
      handleClean();
    }
    setFlagForClean(false);
  }, [flagForClean]);

  return (
    
      <div style={{ maxWidth: 600 }}>
      <form style={{ padding: '20px', border: '1px solid #ccc', borderRadius: '5px', boxShadow: '0 0 5px rgba(0, 0, 0, 0.1)' }}>
  {/* QuizName */}
  <div style={{ marginBottom: '15px', display: 'flex', alignItems: 'center', justifyContent: 'flex-start' }}>
    <label htmlFor="QuizName" style={{ fontWeight: 'bold', marginRight: '10px', width: '100px' }}>Quiz Name:</label>
    <input
      type="text"
      id="QuizName"
      required
      value={quizName}
      onChange={handleInputTitleChange}
      style={{ padding: '5px', borderRadius: '3px', border: '1px solid #ccc', flex: 1 }}
    />
  </div>

  {/* Score */}
  <div style={{ marginBottom: '15px', display: 'flex', alignItems: 'center', justifyContent: 'flex-start' }}>
    <label htmlFor="Score" style={{ fontWeight: 'bold', marginRight: '10px', width: '100px' }}>Score:</label>
    <input
      type="number"
      id="Score"
      required
      value={quizScore}
      onChange={handleInputScoreChange}
      style={{ padding: '5px', borderRadius: '3px', border: '1px solid #ccc', flex: 1 }}
    />
  </div>

  {/* Start Time */}
  <div style={{ marginBottom: '15px', display: 'flex', alignItems: 'center', justifyContent: 'flex-start' }}>
    <label htmlFor="startTimeQuiz" style={{ fontWeight: 'bold', marginRight: '10px', width: '100px' }}>Start Time:</label>
    <input
      type="datetime-local"
      id="startTimeQuiz"
      value={startTimeQuiz}
      onChange={handleStartQuizChange}
      style={{ padding: '5px', borderRadius: '3px', border: '1px solid #ccc', flex: 1 }}
    />
  </div>

  {/* End Time */}
  <div style={{ marginBottom: '20px', display: 'flex', alignItems: 'center', justifyContent: 'flex-start'}}>
    <label htmlFor="endTimeQuiz" style={{ fontWeight: 'bold', marginRight: '10px', width: '100px' }}>End Time:</label>
    <input
      type="datetime-local"
      id="endTimeQuiz"
      value={endTimeQuiz}
      onChange={handleEndQuizChange}
      style={{ padding: '5px', borderRadius: '3px', border: '1px solid #ccc', flex: 1 }}
    />
  </div>
   {/* Render the QuizEdit component */}
   <QuizEdit
          style={{bottom:'0',right:'0'}}
          course_number={courseNumber}
          startTime={startTimeQuiz}
          max_score={quizScore}
          end_time={endTimeQuiz}
          quiz_name={quizName}
          setFlagForClean={setFlagForClean}
        />
</form>



       
      </div>

    
  )
}