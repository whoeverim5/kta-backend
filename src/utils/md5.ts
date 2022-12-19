import crypto from 'crypto';

const md5 = crypto.createHash('md5');
const salt = 'wu-xian-quanNestjsBackend@k!t~a&e*x^a#m';

function randomSalt(): string {
  let rSalt = '';
  for (let i = 0; i <= Math.floor(Math.random() * 6) + 10; i++) {
    rSalt += salt.charAt(Math.floor(Math.random() * salt.length));
  }
  return rSalt;
}

function md5Crypto(password: string): string {
  return md5.update(password).update(randomSalt()).digest('hex');
}
