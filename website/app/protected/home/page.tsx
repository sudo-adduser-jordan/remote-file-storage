import Link from 'next/link';
import styles from './/home.module.css';
import { logout } from '@/lib/requests';

export default function HomePage() {
  return (
    <main className={styles.main}>
      <div className={styles.title}>Remote File Storage</div>

      <div className={styles.container}>
        <div>
          <Link href={'/protected/upload'}>
            <button className={styles.button}> 🡅 Upload</button>
          </Link>
          <Link href={'/protected/download'}>
            <button className={styles.button}> 🡇 Download</button>
          </Link>
        </div>
        <Link href={'/'} className={styles.logout} onClick={logout}>
          Logout
        </Link>
      </div>
    </main>
  );
}

// 🡅 🡇
