import { Module } from '@nestjs/common';
import { FilesController } from './files.controller';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';

@Module({
  imports: [
    ClientsModule.register([
      {
        transport: Transport.GRPC,
        name: 'fileworker',
        options: {
          package: 'fileworker',
          protoPath: join(__dirname, '../../../file.proto'),
          url: `${process.env.GRPC_HOST || 'localhost'}:${process.env.GRPC_PORT || '50051'}`,
        },
      },
    ]),
  ],
  controllers: [FilesController],
})
export class FilesModule {}
