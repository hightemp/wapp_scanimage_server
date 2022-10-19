<div class="main-wrapper">
    <div class="top-panel">Сканирование документов</div>
    <div class="middle-wrapper">
        <div class="left-side">
            <div style="overflow-y:scroll;">
                <?php echo $sScannersList ?>
            </div>
            <form class="left-side-form" action="?m=Main&a=fnActionHTML" target="actions-iframe" method="POST">
                <div>
                    <button type="sumbit" class="btn btn-success" name="action" value="scan">Скан.</button>
                    <button type="sumbit" class="btn btn-danger" name="action" value="delete">Удал.</button>
                    <button type="sumbit" class="btn btn-danger" name="action" value="delete_all">Очис.</button>
                    <button type="sumbit" class="btn btn-info" name="action" value="download">Скач.</button>
                    <button type="button" class="btn btn-primary" name="action" value="refresh" onlcik="window.location.reload()">Обн.</button>
                </div>
                <div style="overflow-y:scroll;">
                    <div class="list-group">
                        <?php foreach ($aScannedFiles as $aScannedFile): ?>
                        <div class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <input type="checkbox" value="<?php echo $aScannedFile[0] ?>" name="images[]" /> 
                                <a href="?m=Main&a=fnImageHTML&file=<?php echo $aScannedFile[0] ?>" target="image-iframe"><h5 class="mb-1"><?php echo $aScannedFile[0] ?></h5></a>
                            </div>
                            <p class="mb-1"><?php echo $aScannedFile[2] ?></p>
                        </div>
                        <?php endforeach ?>
                    </div>
                </div>
                <div>
                    <button type="sumbit" class="btn btn-danger" name="action" value="delete_archives">Удал.</button>
                    <button type="sumbit" class="btn btn-danger" name="action" value="delete_all_archives">Очис.</button>
                </div>
                <div style="overflow-y:scroll;">
                    <div class="list-group">
                        <?php foreach ($aPackedFiles as $aPackedFile): ?>
                        <div class="list-group-item list-group-item-action flex-column align-items-start">
                            <div class="d-flex w-100 justify-content-between">
                                <input type="checkbox" value="<?php echo $aPackedFile[0] ?>" name="archives[]" /> 
                                <a href="<?php echo $aPackedFile[1] ?>" target="_blank"><h5 class="mb-1"><?php echo $aPackedFile[0] ?></h5></a>
                            </div>
                            <p class="mb-1"><?php echo $aPackedFile[2] ?></p>
                        </div>
                        <?php endforeach ?>
                    </div>
                </div>
            </form>
        </div>
        <div class="right-side">
            <iframe src="about:blank" frameborder="0" id="actions-iframe" name="actions-iframe" style="display:none"></iframe>

            <iframe src="?m=Main&a=fnImageHTML" frameborder="0" id="image-iframe" name="image-iframe"></iframe>
        </div>
    </div>
</div>
