import { Observable } from 'rxjs';

export interface ResponseError {
  errorMsg: string;
  errorCode: number;
}

export interface DefaultResponse {
  error: ResponseError;
  msg: string;
}

export interface FileMetadata {
  filename: string;
  encoding: string;
  contentType: string;
  path: string;
  ext: string;
  size: number;
}

export interface File {
  name: string;
  content: Buffer;
  metadata: FileMetadata;
}

export interface Files {
  files: File[];
}

export interface FileNames {
  name: string[];
}

export interface FileWorkerService {
  SaveFile(file: File): Observable<DefaultResponse>;
  SaveFiles(files: Files): Observable<DefaultResponse>;
  DeleteFile(FileNames: FileNames): Observable<DefaultResponse>;
}
