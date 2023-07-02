'use client';
import { useState } from 'react';
import styles from './page.module.css';
import { useRouter } from 'next/navigation';
import { createAccount, login } from '@/lib/requests';

export default function LoginPage() {
  const [error, setError] = useState(false);
  const [submit, setSubmit] = useState('name');
  const [check, setCheck] = useState(false);
  const router = useRouter();

  async function handleSubmit(e: any) {
    e.preventDefault();
    if (submit == 'create') {
      const result = await createAccount(e, router);
      if (result) {
        setError(true);
      }
    }
    if (submit == 'login') {
      const result = await login(e, router);
      if (result) {
        setError(true);
      }
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
            autoComplete='off'
            placeholder='Enter username'
          />
          Password: {error && <>❌</>}
          <input
            className={styles.input}
            name='password'
            autoComplete='off'
            placeholder='Enter password'
          />
          <div className={styles.checkbox}>
            <input name='checkbox' type='checkbox' /> Remember me
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
