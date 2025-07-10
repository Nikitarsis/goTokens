from http.server import HTTPServer, BaseHTTPRequestHandler
import sys

class RequestHandler(BaseHTTPRequestHandler):
  """Кастомный обработчик HTTP-запросов"""
    
  def _log_request(self):
    """Логирование деталей запроса в консоль"""
    # Определяем длину тела запроса
    content_length = self.headers.get('Content-Length')
    body = b''
    if content_length:
      # Читаем тело запроса
      body = self.rfile.read(int(content_length))
        
    sys.stdout.flush()  # Принудительный сброс буфера
        
    # Пытаемся декодировать как UTF-8, иначе выводим байты
    try:
      print(body.decode('utf-8'))
    except UnicodeDecodeError:
      print(body)
    # print('=' * 80)
    
  def _send_response(self):
    """Отправляем стандартный ответ клиенту"""
    self.send_response(200)
    self.send_header('Content-type', 'text/plain')
    self.end_headers()
    self.wfile.write(b'Request logged successfully')
    
  def do_ANY(self):
    """Обработчик для всех HTTP-методов"""
    self._log_request()
    self._send_response()
    
  # Динамически создаем обработчики для всех методов
  for method in ['GET', 'POST', 'PUT', 'DELETE', 'PATCH', 'OPTIONS', 'HEAD']:
    exec(f"def do_{method}(self): self.do_ANY()")

def run_server(port=8000):
  """Запускает HTTP-сервер на указанном порту"""
  server_address = ('', port)
  httpd = HTTPServer(server_address, RequestHandler)
  print(f"Starting HTTP server on port {port}")
  print("Press Ctrl+C to stop...")
  httpd.serve_forever()

if __name__ == '__main__':
  run_server()