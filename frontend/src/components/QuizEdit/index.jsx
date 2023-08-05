import { FormControl, Grid, InputLabel, Menu, MenuItem, Select, Stack, styled, TextField, Typography } from '@mui/material'
import { Button, Modal } from 'react-bootstrap';
import React, {useState } from 'react'


function QuizEdit(props) {
  //console.log(props)
  const token = sessionStorage.getItem('token');
  const [answersOption, setAnswersOption] = useState('');
  const [question, setQuestion] = useState(
    { questionId: '', answerType: '', description: '', score: '',  question_number:'' }
  );
  // all the choice ABCD are string types.
  const [choiceA, setChoiceA] = useState('');
  const [choiceB, setChoiceB] = useState('');
  const [choiceC, setChoiceC] = useState('');
  const [choiceD, setChoiceD] = useState('');
  //const { alert, Alert } = useAlert();
  const [quizId,setQuizId]= useState('');
  const [showModal, setShowModal] = useState(false);
// set the time year, month, day, hour, minute, second
  const formatDateTime = (dateTime) => {
    const date = new Date(dateTime);
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const seconds = String(date.getSeconds()).padStart(2, '0');
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
  };

  const inputAnswerOptionChange = (event) => {
    setAnswersOption(event.target.value);
    //console.log(event.target.value);
    if (question.answerType === 'single' && event.target.value.length > 1)
      alert('Answer is single');
    
    
    // console.log(1,event.target.value);
  }
  


  const handleCloseModal = () => {
    setShowModal(false);
    props.setFlagForClean(true);
    
  };

  const SubmitFunc = () => {
    if (!question.question_number || !Number.isInteger(parseInt(question.question_number))) {
      alert("Question name must be a valid integer.");
      return;
    }
  
    // Check if all required fields are filled in
    if (!question.description || !answersOption || !choiceA || !choiceB || !choiceC || !choiceD || !question.answerType) {
      alert("Please fill in all required fields.");
      return;
    }
  
    // Check if question.score is an integer or not empty
    if (!question.score || !Number.isInteger(parseInt(question.score))) {
      alert("Question score must be a valid integer.");
      return;
    }
    // console.log(question)
    const request = {
      quiz_id: parseInt(quizId),
      question_number: parseInt(question.question_number),
      type: question.answerType === 'single' ? parseInt('1') : parseInt('0'),
      description: question.description,
      score: parseInt(question.score),
      answer: answersOption,
      select_A: choiceA,
      select_B: choiceB,
      select_C: choiceC,
      select_D: choiceD,
    };
    //console.log(request)

    fetch('http://localhost:5005/api/v1/quiz_question', {
      method: 'POST',
      headers: {
        'Authorization': token,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(request),
    })
      .then(response => response.json())
      .then((data) => {
        if (data.status !== 200) {
          alert(data.msg);
        } else {
          // setQuizId(data.data.quiz_id),
          setQuestion({
            questionId: '',
            answerType: '',
            description: '',
            score: '',
            question_number:'',
            answers: [],
            correctAnswerIds: [],
          });
          setChoiceA('');
          setChoiceB('');
          setChoiceC('');
          setChoiceD('');
          setAnswersOption('');
          
          //handleCloseModal();
          alert('Success!');

        }
      })
      .catch((error) => {
        alert("Failed!")
      });
  };

  const handleNavigation = async(props) => {
    if (!props.quiz_name || !props.max_score || !props.startTime || !props.end_time) {
      alert("Quiz name, score, start time, and end time are required.");
      props.setFlagForClean(true);
      return;
    }
    const maxScore = parseInt(props.max_score);
    if (isNaN(maxScore) || maxScore <= 0 || maxScore > 100) {
      alert("Max score must be a number between 1 and 100.");
      props.setFlagForClean(true);
      return;
    }
  
   
    const now = new Date();
    if (props.startTime < now) {
      alert("Start time cannot be in the past.");
      props.setFlagForClean(true);
      return;
    }
  
    if (props.startTime > props.end_time) {
      alert("Start time cannot be greater than end time.");
      props.setFlagForClean(true);
      return;
    }
  
    const timeDifferenceInMinutes = (props.end_time - props.startTime) / (1000 * 60);
    if (timeDifferenceInMinutes < 5) {
      alert("Start time and end time must be at least five minutes apart.");
      props.setFlagForClean(true);
      return;
    }
    setShowModal(true);
    const startTime = formatDateTime(props.startTime);
    const endTime = formatDateTime(props.end_time);

    const request= {
      course_number: parseInt(props.course_number) ,
      quiz_name: props.quiz_name,
      max_score: parseInt(props.max_score) ,
      start_time: startTime ,
      end_time: endTime ,
    };

    fetch('http://127.0.0.1:5005/api/v1/quiz', {
      method: 'POST',
      headers: {
        'Authorization': token,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(request),
    })
      .then(response => response.json())
      .then((data) => {
        if (data.status !== 200) {
          alert(data.msg);
        } else {
          setQuizId(data.data.quiz_id)
          //console.log(data.data.quiz_id)
          alert('Success!');
          //console.log(data.data);

        }
      })
      .catch((error) => {
        alert("Failed!")
      });
  };

return (
    <div>
      {/* Button to open the modal */}
      <Button 
        style={{ marginLeft: '100px' }} variant="primary" color="blue" onClick={() => handleNavigation(props)} 
      >
        Start 
      </Button>

      <Modal show={showModal} onHide={handleCloseModal}>
        <Modal.Header closeButton>
          <Modal.Title>Set Quiz</Modal.Title>
        </Modal.Header>
        <Modal.Body>
        <Stack
            sx={{
              padding: '16px'
            }}
            direction="column"
            justifyContent="start"
            alignItems="stretch"
            spacing={1}
          >
            <TextField label="Question Name" margin="normal" variant="outlined" value={question.question_number} onChange={(e) => setQuestion({ ...question, question_number: e.target.value })} /> <br />
            <TextField label="Question Description" margin="normal" variant="outlined" value={question.description} onChange={(e) => setQuestion({ ...question, description: e.target.value })} /> <br />
            <TextField label="Question Score" type="number" margin="normal" variant="outlined" value={question.score} onChange={(e) => setQuestion({ ...question, score: e.target.value })} /> <br />
            {/* <TextField label="Question Limit(s)" type="number" margin="normal" variant="outlined" value={questionDuration} onChange={(e) => setQuestion({ ...question, duration: e.target.value })} /> <br /> */}
            <FormControl fullWidth>
              <InputLabel id="demo-simple-select-label">Answer Type</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={question.answerType}
                label="Answer Type"
                onChange={(e) => setQuestion({ ...question, answerType: e.target.value })}
              >

                <MenuItem value={'single'}>single</MenuItem>
                <MenuItem value={'multiple'}>multiple</MenuItem>
              </Select>
            </FormControl>
            <Stack
              sx={{
                paddingTop: '36px'
              }}
              direction="row"
              justifyContent="center"
              alignItems="center"
              spacing={2}
            >
              {/* <Typography variant="h5" component="h2">
                Answers
              </Typography> */}
              {/* <Button variant="contained" color="success" 
              disabled={question.answers.length >= 6} 
              onClick={() => addAnswer()}>Add</Button>
              <Button variant="contained" color="error" 
              disabled={question.answers.length <= 2} 
              onClick={() => deleteAnswer()}>Delete</Button> */}
            </Stack>
            <Grid container spacing={2}>
        
          <Grid item xs={3}>
            <TextField
              label="choiceA"
              margin="normal"
              variant="outlined"
              value={choiceA}
              fullWidth
              onChange={(e) => setChoiceA(e.target.value)}
            />
          </Grid>
          <Grid item xs={3}>
            <TextField
              label="choiceB"
              margin="normal"
              variant="outlined"
              value={choiceB}
              fullWidth
              onChange={(e) => setChoiceB(e.target.value)}
            />
          </Grid>
          <Grid item xs={3}>
            <TextField
              label="choiceC"
              margin="normal"
              variant="outlined"
              value={choiceC}
              fullWidth
              onChange={(e) => setChoiceC(e.target.value)}
            />
          </Grid>
          <Grid item xs={3}>
            <TextField
              label="choiceD"
              margin="normal"
              variant="outlined"
              value={choiceD}
              fullWidth
              onChange={(e) => setChoiceD(e.target.value)}
            />
          </Grid>
      </Grid>
            <div style={{ color: 'gray', fontStyle: 'italic' ,marginTop:'10px'}}>
              <p>For single Question, write 'A'</p>
              <p>For multiple Question, if the answer is A and B, write 'AB'</p>
            </div>
            {/* <Button variant="contained" onClick={()=>{handleSubmit()}}></Button> */}
            <div style={{ textAlign: 'left' }}>
              Your answers:
            </div>
            <div style={{ display: 'flex', alignItems: 'center' }}>
              <div style={{ width: '100px', marginRight: '50px' }}>
                <TextField
                  label="Answers"
                  margin="normal"
                  variant="outlined"
                  value={answersOption}
                  onInput={inputAnswerOptionChange}
                  style={{ width: '100%', boxSizing: 'border-box' }}
                />
              </div>
            
            </div>
          </Stack>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleCloseModal}>
            Cancel
          </Button>
          <Button variant="primary" onClick={SubmitFunc}>
            Submit
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
}
export default QuizEdit;