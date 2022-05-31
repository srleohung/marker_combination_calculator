import React, { useState, useEffect } from 'react';
import styles from './Display.module.css';

var progress;

export function Display() {
  const [message, setMessage] = useState("");
  const [status, setStatus] = useState('0');
  const [inProgress, setInProgress] = useState(false);

  const fetchData = async (action) => {
    const param = {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    };
    const response = await fetch("http://localhost:8080/" + action, param);
    const data = await response.json();
    return data.message
  }

  const createNewTask = async () => {
    setInProgress(true)
    let message = await fetchData("new_task")
    setMessage(message)
  }

  const getTaskResult = async () => {
    setInProgress(true)
    let message = await fetchData("get_result")
    setMessage(message)
  }

  const cancelTask = async () => {
    setInProgress(false)
    let message = await fetchData("cancel_task")
    setMessage(message)
  }

  useEffect(() => {
    if (inProgress) {
      progress = setInterval(async () => {
        let message = await fetchData("get_progress")
        setStatus(message);
        if (message == 100) {
          getTaskResult()
          clearInterval(progress);
        }
      }, 500);
    } else {
      clearInterval(progress);
    }
  }, [inProgress]);

  return (
    <div>
      <div className={styles.row}>
        <input
          className={styles.textbox}
          aria-label="Set increment amount"
          value={status}
        />
        <button
          className={styles.button}
          onClick={createNewTask}
        >
          Start Task
        </button>
        <button
          className={styles.asyncButton}
          onClick={cancelTask}
        >
          Cancel Task
        </button>
      </div>
      <div className={styles.row}>
        <span className={styles.value}>{message}</span>
      </div>
    </div>
  );
}
