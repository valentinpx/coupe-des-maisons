<?php
function getPoints($house)
{
    $response = json_decode(file_get_contents("http://cdm.valentinsene.me/api/houses/{$house}/total"));

    return $response->total;
}

$houses = [
    [
        'name' => 'Serdaigle',
        'logo' => './img/Blason_de_Serdaigle.png',
        'points' => getPoints('Serdaigle')
    ],
    [
        'name' => 'Gryffondor',
        'logo' => './img/Blason_de_Gryffondor.png',
        'points' => getPoints('Gryffondor')
    ],
    [
        'name' => 'Poufsouffle',
        'logo' => './img/Blason_de_Poufsouffle.png',
        'points' => getPoints('Poufsouffle')
    ],
    [
        'name' => 'Serpentard',
        'logo' => './img/Blason_de_Serpentard.png',
        'points' => getPoints('Serpentard')
    ]
];

$max = 1;

foreach ($houses as $house)
    if ($house['points'] > $max)
        $max = $house['points'];

?>
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <title>CDM Anubis</title>

    <link href="https://unpkg.com/tailwindcss@^2/dist/tailwind.min.css" rel="stylesheet">
</head>
<body class="bg-black h-screen">
    <div class="p-5 flex justify-around">
        <div class="w-1/3 border-2 border-yellow-500 rounded-lg p-3">
            <h3 class="text-3xl font-medium text-yellow-500 text-center">Coupe des Maisons 2021</h3>
            <p class="m-3 text-lg text-yellow-600">
                Bienvenue à vous, Tek1 !<br>
                Prenez part à des activités organisées par le BDE et remportez des points pour votre maison, la maison gagnante se verra offrir une surprise.<br>
                Une cérémonie sera organisée en fin de piscine.<br>
                Amusez vous bien !<br><br>
                — Le BDE Anubis
            </p>
        </div>
        <div>
            <img class="w-80 mx-auto" src="https://i.imgur.com/3XyvBTu.jpg" alt="">
        </div>
        <div class="w-1/3 border-2 border-yellow-500 rounded-lg p-3">

        </div>
    </div>
    <div class="mx-auto p-5 sm:px-6 lg:px-20 py-5">
        <div class="grid grid-cols-4 gap-10">
            <?php foreach ($houses as $house) { ?>
                <div class="text-center p-5 rounded-lg" style="background-image: url('https://bde-anubis.eu/img/aztec.svg'); background-color: #333">
                    <img class="h-44 mb-8 mx-auto" src="<?php echo $house['logo'] ?>">
                    <div class="relative h-72 w-1/3 mx-auto bg-black rounded-xl">
                        <div class="block absolute bottom-0 w-full <?= $house['points'] == $max ? 'bg-red-600 animate animate-pulse' : 'bg-yellow-500'?> rounded-lg" style="height: <?= $house['points'] / $max * 100 ?>%;">
                        </div>
                    </div>
                    <h3 class="mt-5 text-white text-4xl font-bold"><?= $house['points'] ?> points</h3>
                </div>
            <?php } ?>
        </div>
    </div>
</body>
</html>