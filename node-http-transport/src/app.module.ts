import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { FilesController } from './files/files.controller';
import { FilesService } from './files/files.service';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      envFilePath: process.env.RUN_MODE || '.env.dev',
    }),
  ],
  controllers: [FilesController],
  providers: [FilesService],
})
export class AppModule {}
