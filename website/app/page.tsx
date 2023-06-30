import Link from 'next/link'
import styles from './page.module.css'

export default function Home() {
  return (
    <main className={styles.main}>

      <div className={styles.title}>
        Remote File Storage
      </div>

      <div className={styles.container}>
        Sign In
        <form
          className={styles.form}
        >
          Username:
          <input className={styles.input} placeholder='Enter username' name="username" />
          Password:
          <input className={styles.input} placeholder='Enter password' name="password" />

        {/* <div className={styles.checkbox}>
          <input type="checkbox" /> Remember me
        </div> */}

        </form>

        <div>
      <Link href={'/home'}> 
          <button className={styles.button} >Login</button>
      </Link>
          <button className={styles.button} >Create Account</button>
        </div>

      </div>
    </main>
  )
}
