<div class="main-wrapper">
    <div class="top-panel">Сканирование документов</div>
    <div class="middle-wrapper">
        <div class="left-side">
            <div style="overflow-y:scroll;">
                <?php echo $sScannersList ?>
            </div>
            <form class="left-side-form" action="?m=Main&a=fnActionHTML" target="actions-iframe" method="POST">

                <nav>
                    <div class="nav nav-tabs" id="nav-tab" role="tablist">
                        <!-- <a class="nav-item nav-link" id="nav-scans-tab" data-toggle="tab" href="#nav-scans" role="tab" aria-controls="nav-scans" aria-selected="true">Ск1</a> -->
                        <a class="nav-item nav-link active" id="nav-scans-grid-tab" data-toggle="tab" href="#nav-scans-grid" role="tab" aria-controls="nav-scans-grid" aria-selected="true">Сканы</a>
                        <a class="nav-item nav-link" id="nav-home-tab" data-toggle="tab" href="#nav-home" role="tab" aria-controls="nav-home" aria-selected="true">Архивы</a>
                        <a class="nav-item nav-link" id="nav-profile-tab" data-toggle="tab" href="#nav-profile" role="tab" aria-controls="nav-profile" aria-selected="false">PDF</a>
                        <a class="nav-item nav-link" id="nav-settings-tab" data-toggle="tab" href="#nav-settings" role="tab" aria-controls="nav-settings" aria-selected="false">Настройки</a>
                    </div>
                </nav>
                <div class="tab-content" id="nav-tabContent">
                    <!-- <div class="tab-list tab-pane fade show" id="nav-scans" role="tabpanel" aria-labelledby="nav-scans-tab">
                        <?php require_once("list_scan_tab.php") ?>
                    </div> -->
                    <div class="tab-list tab-pane fade show active" id="nav-scans-grid" role="tabpanel" aria-labelledby="nav-scans-grid-tab">
                        <?php require_once("list_scan_grid_tab.php") ?>
                    </div>
                    <div class="tab-list tab-pane fade show" id="nav-home" role="tabpanel" aria-labelledby="nav-home-tab">
                        <?php require_once("list_packed_files.php") ?>
                    </div>
                    <div class="tab-list tab-pane fade" id="nav-profile" role="tabpanel" aria-labelledby="nav-profile-tab">
                        <?php require_once("list_pdf_files.php") ?>
                    </div>
                    <div class="tab-pane fade" id="nav-settings" role="tabpanel" aria-labelledby="nav-settings-tab">
                        <?php require_once("list_settings.php") ?>
                    </div>
                </div>


            </form>
        </div>
        <div class="right-side">
            <iframe src="<?php echo $sURL ?>" frameborder="0" id="actions-iframe" name="actions-iframe" style="display:none"></iframe>

            <iframe src="?m=Main&a=fnImageHTML" frameborder="0" id="image-iframe" name="image-iframe"></iframe>
        </div>
    </div>
</div>
