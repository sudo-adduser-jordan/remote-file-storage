'use client';
import { useState } from 'react';
import styles from './page.module.css';
import { useRouter } from 'next/navigation';

const LOGIN_PATH = 'http://localhost:8080/login';
const CREATE_PATH = 'http://localhost:8080/register';

export default function LoginPage() {
  const [error, setError] = useState(false);
  const [submit, setSubmit] = useState('name');
  const [check, setCheck] = useState(false);
  const router = useRouter();

  async function login(e: any) {
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);
    console.log(formData);

    const headersList = {
      // Accept: '*/*',
      Authorization: 'Basic dXNlcjE6cGFzcw==',
    };

    const response = await fetch(LOGIN_PATH, {
      method: 'POST',
      headers: headersList,
      body: formData,
    });

    if (response.ok) {
      router.push('/home');
    } else {
      setError(true);
    }
  }

  async function createAccount(e: any) {
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);

    const headersList = {
      Accept: '*/*',
    };

    const response = await fetch(CREATE_PATH, {
      method: 'POST',
      headers: headersList,
      body: formData,
    });
    console.log(response);
    console.log(response.ok);

    if (response.ok) {
      router.push('/home');
    }
  }

  function handleSubmit(e: any) {
    e.preventDefault();
    if (submit == 'login') {
      login(e);
    }
    if (submit == 'create') {
      createAccount(e);
    }
  }

  return (
    <main className={styles.main}>
      <div className={styles.title}>Remote File Storage</div>

      <div className={styles.container}>
        <form
          className={styles.form}
          id='form'
          method='POST'
          onSubmit={handleSubmit}
        >
          Username: {error && <>❌</>}
          <input
            className={styles.input}
            name='username'
            placeholder='Enter username'
          />
          Password: {error && <>❌</>}
          <input
            className={styles.input}
            name='password'
            placeholder='Enter password'
          />
          <div className={styles.checkbox}>
            <input type='checkbox' /> Remember me
          </div>
        </form>
        <div>
          <button
            className={styles.button}
            type='submit'
            form='form'
            onClick={() => {
              setSubmit('login');
            }}
          >
            Login
          </button>
          <button
            className={styles.button}
            type='submit'
            form='form'
            onClick={() => {
              setSubmit('create');
            }}
          >
            Create Account
          </button>
        </div>
      </div>
    </main>
  );
}

// ✔️
// ❌
