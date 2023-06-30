import Link from 'next/link';
import styles from '../page.module.css';

export default function Home() {
  return (
    <main className={styles.main}>
      <div className={styles.title}>Remote File Storage</div>

      <div className={styles.container}>
        <Link href={'/'}>Logout</Link>
      </div>
    </main>
  );
}
