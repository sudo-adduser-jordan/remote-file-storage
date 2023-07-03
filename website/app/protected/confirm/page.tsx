import Link from 'next/link';
import styles from './/confirm.module.css';
import { logout } from '@/lib/requests';

export default function ConfirmPage() {
  return (
    <main className={styles.main}>
      <div className={styles.title}>Remote File Storage</div>

      <div className={styles.container}>
        Success!
        <div>✔️</div>
        <div className={styles.links}>
          <Link href={'/protected/upload'}>
            <button className={styles.button}> 🡅 Upload</button>
          </Link>
          <Link href={'/protected/download'}>
            <button className={styles.button}> 🡇 Download</button>
          </Link>
        </div>
        <div className={styles.logout}>
          <Link href={'/'} onClick={logout}>
            Logout
          </Link>
        </div>
      </div>
    </main>
  );
}
