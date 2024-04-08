import {
  BadRequestException,
  Body,
  Controller,
  Delete,
  HttpCode,
  HttpStatus,
  Inject,
  OnModuleInit,
  Post,
  UploadedFile,
  UploadedFiles,
  UseInterceptors,
} from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { FileInterceptor, FilesInterceptor } from '@nestjs/platform-express';
import { FileWorkerService, DefaultResponse, File } from './files.interface';
import { extname } from 'path';

@Controller('files')
export class FilesController implements OnModuleInit {
  private fileWorkerService: FileWorkerService;

  constructor(@Inject('fileworker') private client: ClientGrpc) {}

  public onModuleInit() {
    this.fileWorkerService =
      this.client.getService<FileWorkerService>('FileWorker');
  }

  @Post('upload')
  @HttpCode(HttpStatus.CREATED)
  @UseInterceptors(FileInterceptor('file'))
  public async saveFile(
    @UploadedFile() file: Express.Multer.File,
  ): Promise<DefaultResponse> {
    if (typeof file === 'undefined' || !file) {
      new BadRequestException('file should be attached');
    }

    return this.fileWorkerService
      .SaveFile({
        name: file.originalname,
        content: file.buffer,
        metadata: {
          filename: file.originalname,
          contentType: file.mimetype,
          encoding: file.encoding,
          ext: extname(file.originalname),
          path: file.path,
          size: file.size,
        },
      })
      .toPromise();
  }

  @Post('upload-many')
  @HttpCode(HttpStatus.CREATED)
  @UseInterceptors(FilesInterceptor('files'))
  public async saveFiles(
    @UploadedFiles() files: Express.Multer.File[],
  ): Promise<DefaultResponse> {
    if (!files || !files.length) {
      throw new BadRequestException('Files should be attached');
    }

    const filesToSave = files.map<File>((file) => ({
      name: file.originalname,
      content: file.buffer,
      metadata: {
        filename: file.originalname,
        contentType: file.mimetype,
        encoding: file.encoding,
        ext: extname(file.originalname),
        path: file.path,
        size: file.size,
      },
    }));

    return this.fileWorkerService.SaveFiles({ files: filesToSave }).toPromise();
  }

  @Delete()
  public async deleteFiles(
    @Body() { files }: { files: string[] },
  ): Promise<DefaultResponse> {
    if (!files || !files.length) {
      throw new BadRequestException('Files should be attached');
    }

    return this.fileWorkerService.DeleteFile({ name: files }).toPromise();
  }
}
