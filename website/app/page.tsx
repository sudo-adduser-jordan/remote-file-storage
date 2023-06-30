'use client';
import Link from 'next/link';
import styles from './page.module.css';

const login_path = 'http://localhost:8080/login';

export default function Login() {
  async function handleSubmit(e: any) {
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);

    const headersList = {
      Accept: '*/*',
      Authorization: 'Basic dXNlcjE6cGFzcw==',
    };

    const response = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: headersList,
      body: formData,
    });

    const data = await response.text();
    console.log(data);
  }

  return (
    <main className={styles.main}>
      <div className={styles.title}>Remote File Storage</div>

      <div className={styles.container}>
        Sign In
        <form className={styles.form} method='post' onSubmit={handleSubmit}>
          Username:
          <input
            className={styles.input}
            name='username'
            placeholder='Enter username'
          />
          Password:
          <input
            className={styles.input}
            name='password'
            placeholder='Enter password'
          />
          {/* <div className={styles.checkbox}>
          <input type="checkbox" /> Remember me
        </div> */}
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
