<?php 
use Hightemp\WappSnanimageServer\Project;
?>
<div style="display:flex">
    <button type="sumbit" class="btn btn-lg btn-success" name="action" value="scan">Сканировать</button>
    <button type="button" class="btn btn-lg btn-primary" name="action" value="refresh" onlcik="window.location.reload()">Обновить</button>

    <div class="dropdown">
        <button class="btn btn-lg btn-secondary dropdown-toggle" type="button" id="dropdownMenu2" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
            Действия
        </button>
        <div class="dropdown-menu" aria-labelledby="dropdownMenu2">
            <button type="sumbit" class="dropdown-item" name="action" value="rename_images">Переименовать</button>
            <button type="sumbit" class="dropdown-item" name="action" value="delete">Удалить выделенное</button>
            <button type="sumbit" class="dropdown-item" name="action" value="delete_all">Удалить все</button>
            <button type="sumbit" class="dropdown-item" name="action" value="download">Архивировать выделенное</button>
            <button type="sumbit" class="dropdown-item" name="action" value="download_all">Архивировать все</button>
            <button type="sumbit" class="dropdown-item" name="action" value="convert_pdf">Конвертировать в PDF выделенное</button>
            <button type="sumbit" class="dropdown-item" name="action" value="convert_pdf_all">Конвертировать в PDF все</button>
        </div>
    </div>
</div>
<div style="overflow-y:scroll;">
    <div class="" style="display: grid; grid-template-columns: repeat(3, 1fr);">
        <?php foreach ($aScannedFiles as $aScannedFile): ?>
            <div class="card" style="">
                <?php $sPath = Project::fnGetImageRelPath($aScannedFile[0]); ?>
                <a href="?m=Main&a=fnImageHTML&file=<?php echo $aScannedFile[0] ?>" target="image-iframe">
                    <img class="card-img-top" src="<?php echo $sPath ?>" alt="" loading="lazy">
                    </a>
                <div class="card-body">
                    <label>
                        <input type="checkbox" value="<?php echo $aScannedFile[0] ?>" name="images[]" /> 
                        <?php echo $aScannedFile[0] ?>
                        <p class="mb-1"><?php echo $aScannedFile[2] ?></p>
                    </label>
                </div>
            </div>
        <?php endforeach ?>
    </div>
</div>