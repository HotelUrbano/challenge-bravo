import { Request, Response } from 'express';
import { IDeleteCoinRequestDTO } from './DeleteCoinDTO';
import { DeleteCoinUseCase } from './DeleteCoinUseCase';

export class DeleteCoinController {
  constructor(private deleteCoinUseCase: DeleteCoinUseCase) {}

  async handle(request: Request, response: Response): Promise<Response> {
    const { uid }: IDeleteCoinRequestDTO = request.body;

    try {
      await this.deleteCoinUseCase.execute({
        uid,
      });

      return response.status(201).send();
    } catch (err) {
      return response.status(400).json({
        message: err.message || 'Unexpected error.',
      });
    }
  }
}
