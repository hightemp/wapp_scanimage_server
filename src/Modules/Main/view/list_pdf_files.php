<div style="display:flex">
    <button type="button" class="btn btn-lg btn-primary" name="action" value="refresh" onlcik="window.location.reload()">Обновить</button>

    <div class="dropdown">
        <button class="btn btn-lg btn-secondary dropdown-toggle" type="button" id="dropdownMenu2" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
            Действия
        </button>
        <div class="dropdown-menu" aria-labelledby="dropdownMenu2">
            <button type="sumbit" class="dropdown-item" name="action" value="rename_pdf">Переименовать</button>
            <button type="sumbit" class="dropdown-item" name="action" value="delete_pdf">Удалить выделенное</button>
            <button type="sumbit" class="dropdown-item" name="action" value="delete_all_pdf">Удалить все</button>
        </div>
    </div>
</div>
<div style="overflow-y:scroll;">
    <div class="list-group">
        <?php foreach ($aPDFFiles as $aPDFFile): ?>
        <div class="list-group-item list-group-item-action flex-column align-items-start">
            <div class="d-flex w-100 justify-content-between">
                <input type="checkbox" value="<?php echo $aPDFFile[0] ?>" name="pdf[]" /> 
                <a href="<?php echo $aPDFFile[1] ?>" target="image-iframe"><h5 class="mb-1"><?php echo $aPDFFile[0] ?></h5></a>
            </div>
            <p class="mb-1"><?php echo $aPDFFile[2] ?></p>
        </div>
        <?php endforeach ?>
    </div>
</div>