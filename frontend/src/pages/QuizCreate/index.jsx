import { Button, FormControl, Grid, InputLabel, Menu, MenuItem, Select, Stack, styled, TextField, Typography } from '@mui/material'
import React, { useEffect } from 'react'
import { useOutletContext } from 'react-router-dom';
import { useAlert } from '../../hooks/useAlert';
import { Collapse } from "antd";
import './style.css'
const Input = styled('input')({
  display: 'none',
});
export default function QuizCreate() {
  // const { question, setQuestion } = useOutletContext();
  const { alert, Alert } = useAlert();

  const question = [{
    id: 1,
    content: {
      answerType: 'single',
      name: 'question abc',
      description: "xxx",
      score: 89,
      duration: 2,
      answers: [{ idAns: 101, content: '' }]
    }
  },
  {
    id: 2,
    content: {
      answerType: 'single',
      name: 'question abc',
      description: "xxxx",
      score: 88,
      duration: 2,
      answers: [{ idAns: 100, content: '' }]
    }
  }]

  const addAnswer = (num = 1) => {
    for (let i = 0; i < num; i++) {
      question[1].content.answers.push(i
      );
    }
  }

  const deleteAnswer = () => {
    question[1].content.answers.pop();
  }
  return (
    <div className="createQuiz" >
      <Collapse defaultActiveKey={['1']} className='panel'>
        {Alert}
        {
          question && (
            <div className="scrollable-content">
              <Stack
                sx={{
                  padding: '16px'
                }}
                direction="column"
                justifyContent="start"
                alignItems="stretch"
                spacing={1}
              >
                <TextField label="Question Name" margin="normal" variant="outlined" value={question.name} /> <br />
                <TextField label="Question Description" margin="normal" variant="outlined" value={question.description} /> <br />
                <TextField label="Question Score" type="number" margin="normal" variant="outlined" value={question.score} /> <br />
                <TextField label="Question Limit(s)" type="number" margin="normal" variant="outlined" value={question.duration} /> <br />
                <FormControl fullWidth>
                  <InputLabel id="demo-simple-select-label">Answer Type</InputLabel>
                  <Select
                    labelId="demo-simple-select-label"
                    id="demo-simple-select"
                    value={question.answerType}
                    label="Answer Type"
                  // onChange={(e) => setQuestion({ ...question, answerType: e.target.value })}
                  >
                    <MenuItem value={'single'}>single</MenuItem>
                    <MenuItem value={'multiple'}>multiple</MenuItem>
                    <MenuItem value={'short-answer'}>short answer</MenuItem>
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
                  <Typography variant="h5" component="h2">
                    Answers
                  </Typography>
                  <Button variant="contained" color="success" onClick={() => addAnswer()}>Add</Button>
                  <Button variant="contained" color="error" onClick={() => deleteAnswer()}>Delete</Button>
                </Stack>
                <div className="scrollable-answers">
                  <Grid alignItems="center" container spacing={2}>
                    {
                      question[1].content.answers.map((item, index) => (
                        <Grid key={item.id}>
                          <Stack
                            sx={{
                              padding: '16px'
                            }}
                            direction="column"
                            justifyContent="start"
                            alignItems="stretch"
                            spacing={1}
                          >
                            <TextField label={`Answer ${index + 1}`} margin="normal" variant="outlined"
                              value={item.content} />
                            <Button variant="contained" >

                            </Button>
                          </Stack>
                        </Grid>
                      ))
                    }

                  </Grid>
                </div>
              </Stack>
            </div>
          )
        }
      </Collapse >
    </div >
  )
}

