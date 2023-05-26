<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Exif Results</title>
    <script crossorigin src="https://unpkg.com/react@17/umd/react.production.min.js"></script>
    <script crossorigin src="https://unpkg.com/react-dom@17/umd/react-dom.production.min.js"></script>
    <script src="https://unpkg.com/@babel/standalone/babel.min.js"></script>
    <script type="text/babel">
        const columns = [%s]
        const files = %s
        const ReactAppFromCDN = ()=>{
            return (
                <>
                    <h1 className="title">Results</h1>
                    <table>
                        <tr>
                            {columns.map((name) => {
                                return <th>{name}</th>
                            })}
                        </tr>
                        {
                            files.map((file) => {
                                return <tr>
                                    <td><a href={`https://www.google.com/maps/search/?api=1&query=${file.latitude},${file.longitude}`} target="_blank">{file.filename}</a></td>
                                    <td>{file.latitude}</td>
                                    <td>{file.longitude}</td>
                                </tr>
                            })
                        }
                    </table>
                </>
            )
        }
        ReactDOM.render(<ReactAppFromCDN />, document.querySelector('#root'));
    </script>
</head>
<body>
    <div id="root"></div>
</body>
</html>