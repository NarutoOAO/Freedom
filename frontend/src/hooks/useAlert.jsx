import React from 'react';
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export function useAlert () {
  const [open, setOpen] = React.useState(false);
  const [severity, setSeverity] = React.useState('success');
  const [message, setMessage] = React.useState('')

  const handleClose = () => {
    setOpen(false);
  }

  // open an alert
  const openAlert = (param) => {
    setSeverity(param.severity || 'success');
    setMessage(param.msg || '');
    setOpen(true);
  }

  // open the error Alert
  const error = (msg) => {
    openAlert({
      severity: 'error',
      msg
    })
  }

  // open info Alert
  const info = (msg) => {
    openAlert({
      severity: 'info',
      msg
    })
  }

  // open warning Alert
  const warning = (msg) => {
    openAlert({
      severity: 'warning',
      msg
    })
  }

  // open success Alert
  const success = (msg) => {
    openAlert({
      severity: 'success',
      msg
    })
  }

  return {
    Alert: (
      <Snackbar open={open} anchorOrigin={{ vertical: 'top', horizontal: 'center' }} autoHideDuration={1200} onClose={handleClose}>
        <Alert onClose={handleClose} severity={severity} sx={{ width: '100%' }}>
          {message}
        </Alert>
      </Snackbar>
    ),
    alert: {
      error,
      info,
      warning,
      success,
    }
  };
}
