'use client';
import { useState } from 'react';
import styles from './page.module.css';
import { useRouter } from 'next/navigation';

const PATH = 'http://localhost:8080/login';

export default function LoginPage() {
  const [error, setError] = useState(false);
  const [check, setCheck] = useState(false);
  const router = useRouter();

  async function login(e: any) {
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);

    const headersList = {
      // Accept: '*/*',
      Authorization: 'Basic dXNlcjE6cGFzcw==',
    };

    const response = await fetch(PATH, {
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

  return (
    <main className={styles.main}>
      <div className={styles.title}>Remote File Storage</div>

      <div className={styles.container}>
        Sign In
        <form className={styles.form} method='post' onSubmit={login}>
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
          <div>
            <button className={styles.button} type='submit'>
              Login
            </button>
            <button className={styles.button}>Create Account</button>
          </div>
        </form>
      </div>
    </main>
  );
}

// ✔️
// ❌
