import { Injectable, OnModuleInit } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { Client, ClientGrpc, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { Observable } from 'rxjs';

interface IFileService {
  saveFile(file: { content: Buffer; name: string }): Observable<{}>;
}

@Injectable()
export class FilesService implements OnModuleInit {
  constructor() {}

  @Client({
    transport: Transport.GRPC,
    options: {
      package: 'fileworker',
      protoPath: join(__dirname, '../../../file.proto'),
      url:
        `${process.env.GRPC_HOST}:${process.env.GRPC_PORT}` ||
        'localhost:50051',
    },
  })
  private readonly client: ClientGrpc;

  private fileService: IFileService;

  public onModuleInit() {
    this.fileService = this.client.getService<IFileService>('FileWorker');
  }

  public saveFile(fileBuffer: Buffer, name: string): Observable<{}> {
    return this.fileService.saveFile({ content: fileBuffer, name });
  }
}
