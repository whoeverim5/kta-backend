import {
  Entity,
  Column,
  PrimaryGeneratedColumn,
  CreateDateColumn,
  UpdateDateColumn,
  DeleteDateColumn,
  Generated,
} from 'typeorm';
import {
  Varchar32,
  Varchar64,
  Varchar255,
  DateTime,
} from 'src/config/database';

enum UserRole {
  ADMIN = 'admin',
  STUDENT = 'student',
  TEACHER = 'teacher',
}

@Entity()
export class User {
  @PrimaryGeneratedColumn()
  id: number;

  @Column(Varchar32)
  account: string;

  @Column(Varchar64)
  password: string;

  @Column({ type: 'enum', enum: UserRole, default: UserRole.STUDENT })
  role: UserRole;

  @Generated('uuid')
  @Column({
    ...Varchar255,
    name: 'user_id',
    comment: '用户唯一ID',
    nullable: false,
    unique: true,
    update: false,
  })
  userId: string;

  @Column({ type: 'tinyint', nullable: true })
  gender: number;

  @Column({ type: 'tinyint', nullable: true })
  age: number;

  @Column(Varchar255)
  email: string;

  @Column(Varchar255)
  salt: string;

  @CreateDateColumn()
  @Column({ ...DateTime, name: 'create_at' })
  createAt: Date;

  @UpdateDateColumn()
  @Column({ ...DateTime, name: 'update_at' })
  updateAt: Date;

  @DeleteDateColumn()
  @Column({ ...DateTime, name: 'delete_at', nullable: true })
  deleteAt: Date;
}
