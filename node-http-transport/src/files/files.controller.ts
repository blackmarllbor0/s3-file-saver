import {
  BadRequestException,
  Controller,
  Post,
  UploadedFile,
  UseInterceptors,
} from '@nestjs/common';
import { FileInterceptor } from '@nestjs/platform-express';
import { FilesService } from './files.service';

@Controller('files')
export class FilesController {
  constructor(private readonly fileService: FilesService) {}

  @Post()
  @UseInterceptors(FileInterceptor('file'))
  public async saveFile(
    @UploadedFile() file: Express.Multer.File,
  ): Promise<{ msg: string }> {
    if (!file) {
      new BadRequestException('file should be attached');
    }

    this.fileService.saveFile(file.buffer, file.originalname);
    return { msg: 'File successfully saved' };
  }
}
