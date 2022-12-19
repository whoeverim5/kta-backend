import { ConfigModule, ConfigService } from '@nestjs/config';
import {
  TypeOrmModuleAsyncOptions,
  TypeOrmModuleOptions,
} from '@nestjs/typeorm';
import { ColumnOptions } from 'typeorm';

const MySQL: TypeOrmModuleAsyncOptions = {
  imports: [ConfigModule],
  inject: [ConfigService],
  useFactory: async (
    configuration: ConfigService
  ): Promise<TypeOrmModuleOptions> => ({
    type: configuration.get<string>('DATABASE_TYPE', 'mysql') as 'mysql',
    username: configuration.get<string>('DATABASE_USERNAME', 'root'),
    password: configuration.get<string>('DATABASE_PASSWORD', 'womensange3'),
    host: configuration.get<string>('DATABASE_HOST', 'localhost'),
    port: configuration.get<number>('DATABASE_PORT', 3306),
    database: configuration.get<string>('DATABASE_DATABASE', 'exam'),
    synchronize: configuration.get<boolean>('DATABASE_SYNCHRONIZE', true),
    retryDelay: configuration.get<number>('DATABASE_RETRYDELAY', 500),
    retryAttempts: configuration.get<number>('DATABASE_RETRYATTEMPTS', 10),
    autoLoadEntities: configuration.get<boolean>(
      'DATABASE_AUTOLOADENTITIES',
      true
    ),
  }),
};

const Varchar255: ColumnOptions = {
  type: 'varchar',
  length: 255,
};

const Varchar32: ColumnOptions = {
  type: 'varchar',
  length: 32,
};

const Varchar64: ColumnOptions = {
  type: 'varchar',
  length: 64,
};

const TimeStamp: ColumnOptions = {
  type: 'timestamp',
};

const DateTime: ColumnOptions = {
  type: 'datetime',
};

export { MySQL, Varchar32, Varchar64, Varchar255, TimeStamp, DateTime };
